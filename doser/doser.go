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
need to wait to get back to a normal level. What you do with this information,
I have no idea!

Assumptions: 
This is an oral route of administration 
A constant dosage
A ~consistent timing interval for dosages
A linear superposition for dosages (eg secretion isn't related to concentration) 
*/

//TODO: Calc peak/trough variation in steady state dose withing timing 
//TODO: Calc current value at point in time
//TODO: Calc time until under certain value of mass

/*
dosage_amount := 300.0 //example 300mg
dosage_interval := 12 //example 2/day
t_0 := 75 //example first dose n hours in past
verbose := true
*/

func main() {
    //note max is 15hr
    t_1_2_min := 12 //min range of half life in hours
    k_e_min := math.Ln2/float64(t_1_2_min)//elimination rate constant min

    args := os.Args

    if len(args) != 4 {
        log.Fatal("Expect: doser [dose amount in mg] [dose frequency in hr] [t0 for initial dose in hours ago]")
    }

    dosage_amount, err := strconv.ParseFloat(args[1], 64)
    if err != nil {
        log.Printf("Couldn't convert %s to a float", args[1])
    }
    dosage_interval, err := strconv.ParseInt(args[2], 10, 0)
    if err != nil {
        log.Printf("Couldn't convert %s to an integer", args[2])
    }
    t_0, err := strconv.ParseInt(args[3], 10, 0)
    if err != nil {
        log.Printf("Couldn't convert %s to an integer ", args[3])
    }
    verbose := true

    remaining := 0.0
    var d int64
    for d = 0; d < (t_0 / dosage_interval) + 1; d++ {
        //iterate and sum the doses starting with the first at t_0
        //initial amount * exponential decay function 
        remainder_of_dose := dosage_amount * math.Exp(-k_e_min * (float64(t_0 - d * dosage_interval)))
        if verbose {
            fmt.Printf("Dose from %d hours ago has %f remaining\n", t_0 - d * dosage_interval, remainder_of_dose)
        }
        remaining += remainder_of_dose
    }
    fmt.Printf("Amount left in body: %f\n", remaining)
}
