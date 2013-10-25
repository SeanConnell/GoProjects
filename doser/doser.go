package main

import (
    "fmt"
    "os"
    "math"
    "log"
    "strconv"
)

/*
This program does not represent medical advice in any way shape or form.
User assumes all risk/consequences of this information.

Circulating Modafinil Calculator!
Figure out how much you have in your system right now, and how long you'll
need to wait to get back to a normal level. 
What you'll do with this information, I have no idea!

Assumptions: 
    An intravenous route of administration (mostly because I don't have a good model for oral -- yet)
    A constant dosage
    A ~consistent timing interval for dosages
    A linear superposition for dosages (eg secretion isn't related to concentration) 
    A half life of 13.5 hours, which is the average of the range given for modafinil of 12-15 hours
*/

//TODO Add these features maybe?
//Calc peak/trough variation in steady state dose withing timing 
//Calc time until under certain value of mass

func concentration(dose int, interval int, t_0 int, k_e float64)(remaining float64) {
    for d := 0; d < (t_0 / interval) + 1; d++ {
        //iterate and sum the doses starting with the first at t_0
        //initial amount * exponential decay function 
        remaining += float64(dose) * math.Exp(-k_e * (float64(t_0 - d * interval)))
    }
    return
}

func parseInt(s string)(int){
    i, _ := strconv.ParseInt(s, 10, 0)
    return int(i)
}

func main() {
    t_1_2 := (15+12/2.0) //note: range for the half life is 12-15. we'll assume avg of those two 
    k_e := math.Ln2/t_1_2//elimination rate constant

    args := os.Args
    if len(args) != 4 {
        log.Fatal("Expect: doser [dose amount in mg] [dose frequency in hr] [t0 for initial dose in hours ago]")
    }

    //trust the user. Scary thought.
    dose := parseInt(args[1])
    interval := parseInt(args[2])
    t_0 := parseInt(args[3])

    fmt.Println("hours_from_initial_dose concentration_mg")
    for t := 0; t < t_0; t += 1{
        fmt.Printf("%d %f\n", t, concentration(dose, interval, t, k_e))
    }
}
