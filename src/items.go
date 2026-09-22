package main 

import ( 
    "fmt" 

    "time" 
) 

func findItemIndex(inventaire []string, nom string) int {
    for i, item := range inventaire {
        if item == nom {
            return i
        }
    }
    return -1
}

func takePot(c *Character) { 
    i := findItemIndex(c.Inventaire, "potion") 

    switch { 
    case i == -1: 
        fmt.Println("Sergent : « Ta besace est VIDE, pas de potion, pas de soin ! »") 
    case c.Pv >= c.Pvmax: 
        fmt.Println("Sergent : « PV au max, on ne gâche pas une potion ! »") 
    default: 
        c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...) 
        c.Pv += 50 
        if c.Pv > c.Pvmax { 
            c.Pv = c.Pvmax 
        } 
        fmt.Printf("Sergent : « Potion avalée ! PV : %d/%d »\n", c.Pv, c.Pvmax) 
    } 
} 

func isDead(c *Character) bool { 
    if c.Pv <= 0 { 
        fmt.Printf("%s est mort !\n", c.Inventaire) 
        c.Pv = c.Pvmax / 2 
        fmt.Printf("%s ressuscite avec %d/%d PV\n", c.Inventaire, c.Pv, c.Pvmax) 
        return true 
    } 
    return false 
} 

func poisonPot(c *Character) { 
    fmt.Printf("%s boit la Potion de poison...\n", c.Inventaire) 
 
    const ( 
        degatsParSeconde = 10 
        dureeSecondes    = 3 
    ) 

    for tick := 1; tick <= dureeSecondes; tick++ { 
        time.Sleep(1 * time.Second) 

        c.Pv -= degatsParSeconde 
        if c.Pv < 0 { 
            c.Pv = 0 
        } 

        fmt.Printf("[Poison %d/3s] %s : %d/%d PV\n", tick, c.Inventaire, c.Pv, c.Pvmax) 

        if c.Pv == 0 { 
            isDead(c) 
            return 
        } 
    } 
    isDead(c) 
} 

 