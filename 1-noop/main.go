package main

import "gopkg.in/distil.v1"

func main() {
	// Get a handle to BTrDB and Mongo. go-distil is implemented as a library
	// so there is no other distillate service to connect to
	ds := distil.NewDISTIL(distil.FromEnvVars())

	// Clearly you could have more advanced logic here, but this serves as
	// a good example. This would register the distillate for L1ANG of
	// every PMU that has a nonempty L1MAG stream.
	// for _, path := range ds.ListExistingUpmuPaths() {
	// 	trimPath := strings.TrimPrefix(path, "/upmu/")
	// 	instance := &NopDistiller{}
	// 	registration := &distil.Registration{
	// 		Instance:   instance,
	// 		UniqueName: "noop_" + strings.Replace(trimPath, "/", "_", -1),
	// 		InputPaths: []string{path},
	// 		OutputPaths: []string{path},
	// 	}
	// 	ds.RegisterDistillate(registration)
	// }

	// Construct an instance of your distillate. If you had parameters for
	// the distillate you would maybe have a custom constructor. You could
	// also load the parameters from a config file, or some heuristic
	// algorithm, which we will show in the next few examples
	instance := &NopDistiller{}

	// Now we add this distillate to the DISTIL engine. If you add multiple
	// distillates, they will all get computed in parallel.
	ds.RegisterDistillate(&distil.Registration{
		// The class that implements your algorithm
		Instance: instance,
		// A unique name FOR THIS INSTANCE of the distillate. If you
		// are autogenerating distillates, take care to never produce
		// the same name here. We would normally use a UUID but opted
		// for this so as to be more human friendly. If the program
		// is restarted, this is how it knows where to pick up from.
		UniqueName: "aminy_multiplier_distillate_0",
		// These are inputs to the distillate that will be loaded
		// and presented to Process()
		InputPaths: []string{"/REFSET/LBNL/a6_bus1/L2MAG"},
		// These are the output paths for the distillate. They must
		// also be strictly unique.
		OutputPaths: []string{"/aostfeld/a6_bus1/multiplied"},
	})

	//Now we tell the DISTIL library to keep all the registered distillates
	//up to date. The program will not exit.
	ds.StartEngine()
}
