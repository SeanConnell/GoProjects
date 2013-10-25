## Doser: A blood level concentration calculator
This program is not medical advice, not intended to diagnose, nor provide guidance, or to replace a doctor's advice.

Now that that's out of the way: Doser is intended as an educational tool to explore blood concentrations for various drugs that exhibit exponential elimination from the body. 
Right now it's hard coded to do modafinal blood concentrations.

### Why would I use this?

+ As an exploratory tool to see how much of your prescription you have floating around in your blood at any one time
+ To see how far you drift from peak blood level to trough
+ To examine how long it takes you to reach "steady state" of concentrations from zero.

### Installation/Dependencies

I'm assuming if you're running shell scripts you know a thing or two about how to computer. You'll need go, as well as optionally gnuplot.

http://golang.org/doc/install

http://www.gnuplot.info/download.html

Both of these are available from source as well as in many package repositories. It should be easy to get them installed on almost any platform. Even windows.

The examples below treat this file as a script with ```go run```. You can obviously build it properly and keep it around forever; your own private binary.

### Usage

#### Generating concentration data
Invoke the program on the command line with arguments of the form: ```sh go run doser.go [dose amount in mg] [dose frequency in hr] [t0 for initial dose in hours ago] ```

It will produce something along the lines of
```
hours_from_initial_dose concentration_mg
0 200.000000
1 193.506356
2 187.223548
3 181.144733
4 175.263286
5 169.572798
6 164.067071
7 158.740105
   (etc)
```

This is a representation of the time from first dose in hours (in the first column), and blood concentration at that time in mg. (second column)
#### Visualization with GNUPlot
Save the output of a data run however you'd like. The simplest is with output redirection:

```sh
go run doser.go 200 12 120 > data_file.txt
```

Plot the data file with some basic settings in GNUPlot 

```sh    
echo 'set term png; set xlabel "concentration mg"; set ylabel "hours from first dose"; plot "data_file.txt" using 1:2 with lines' | gnuplot > concentration.png
```

Examine the output with your favorite image viewer

```sh
display concentration.png
```
