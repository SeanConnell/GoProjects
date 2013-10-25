# Doser Modafinil blood level concentration calculator
This program is not medical advice, intended to diagnose, provide guidance, or replace a doctor's advice.

Now that that's out of the way: Doser is intended as a simple blood concentration calculator for various drugs. Right now it's hard coded to do modafinal blood concentrations.

# Usage

## Generating concentration data
Invoke the program on the command line with arguments of the form: 
    doser [dose amount in mg] [dose frequency in hr] [t0 for initial dose in hours ago]

## Visualization with GNUPlot
1. Save the output of a data run however you'd like. The simplest is with output redirection:
    doser 200 12 120 > data_file.txt
2. Plot the data file with some basic settings in GNUPlot 
    echo 'set term png; set xlabel "concentration mg"; set ylabel "hours from first dose"; plot "data_file.txt" using 1:2 with lines' | gnuplot > concentration.png
3. Examine the output with your favorite image viewer
    display concentration.png
