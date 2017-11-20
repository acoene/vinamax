//This example investigates the speed of roation of a single particle when the anisotropy axis, is perpendicual to the external field and m is between both, without temperature
package main

import (
	. "github.com/JLeliaert/vinamax"
	//"fmt"
)

func main() {

	//Defines the world at location 0,0,0 and with a side of 2 m	
	World(0,0,0,2)
	//Adds a cube to the center of the world with side 2 m
	test := Cube{S:2}

	//additionally calculate Brownian rotation
	BrownianRotation = true
	//this requires a randomn number for the anisotropy dynamics
	Setrandomseed_anis(2)
	//set viscosity environement test (e.g. water 1mPas)
	test.Setviscosity(1.e-3)

	//the particle have radius 25 nm and a hydrodynamic radius (core + coating together) of 30 nm (tauN = 4*10^58 s TauB = 80 µs)
	Particle_radius(12.5e-9)
	Particle_radius_h(15.e-9)

	//Adds a single particles with radii defined above to the cube with viscosity 1 mPas
	test.Addparticles(1)
	//fmt.Println("particles added")
	//Don't calculate the demagnetising field
	Demag=false

	//saturation magnetisation 860 000 A/m
	Msat (860e3)
	
	//apply an external field along the z direction of 10mT
	//B_ext can be an arbitrary function of time
	B_ext = func(t float64) (float64, float64, float64) { return 0, 0, 1e-3}

	Adaptivestep = true
	//timestep : 2ps
	Dt = 2e-12
	//initialise time at zero
	T = 0.
	//no temperature
	Temp = 0.
	//The Gilbert damping constant =0.1
	Alpha = 0.1
	//anisotropy constant=10 000 J/m**3
	Ku1 = 10000. 

	//anisotropy axis is in y-direction for the particle (perpendicular to Bext)
	Anisotropy_axis(0,1,0)
	
	//initialise the magnetisation bewteen u_anis and h_ext
	M_uniform(0,1,1)
	
	//curently anisodynamics only works with Euler
	Setsolver("dopri")
	
	//output u_anis of single particle (works only for single particle)
	Tableadd("u_anis")

	//write output every 0.1 µs 
	Output(0.01e-6)
	
	//fmt.Println("dt:   ", Dt)
	//run for 3*tauB 
	Run(1e-4)
}