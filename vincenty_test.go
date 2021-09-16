package vincenty

import "testing"

func TestVincentyDistance(t *testing.T) {
	// distance NY-SD: 3915340.577 m
	// distance SD-EP: 1011300.217 m
	// distance EP-SL: 1663833.491 m
	// distance SL-NY: 1406519.972 m
	var newYork = Coord{40.7128,74.0060}
	var sanDiego = Coord{32.7157,117.1611}
	var elPaso = Coord{31.7619, 106.4850}
	var stLouis = Coord{38.6270,90.1994}

	var miles, km float64
	var ok bool

	miles, km, ok = VincentyDistance(newYork,sanDiego)
	if int(miles) != 2432 || int64(km) != 3915 || ok == false {
		t.Errorf("Computed values: %v %10f %v\n", miles,km,ok)
		t.Errorf("In correct computation between New York and San Diego: %v %v %v\n", int(miles), int64(km), ok)
	}

	miles, km, ok = VincentyDistance(sanDiego,elPaso)
	if int(miles) != 628 || int64(km) != 1011 || ok == false {
		t.Errorf("Computed values: %v %10f %v\n", miles,km,ok)
		t.Errorf("In correct computation between San Diego and El Paso: %v %v %v\n", int(miles), int64(km), ok)
	}

	miles, km, ok = VincentyDistance(elPaso,stLouis)
	if int(miles) != 1033 || int(km) != 1663 || ok == false {
		t.Errorf("Computed values: %v %10f %v\n", miles,km,ok)
		t.Errorf("In correct computation between El Paso and St. Louis: %v %v %v\n", int(miles), int64(km), ok)
	}

	miles, km, ok = VincentyDistance(stLouis,newYork)
	if int(miles) != 873 || int(km) != 1406 || ok == false {
		t.Errorf("Computed values: %v %10f %v\n", miles,km,ok)
		t.Errorf("In correct computation between St. Louis and New York: %v %v %v\n", int(miles), int64(km), ok)
	}

	miles, km, ok = VincentyDistance(newYork,newYork)
	if int(miles) != 0 || int(km) !=  0 || ok == false {
		t.Errorf("Computed values: %v %10f %v\n", miles,km,ok)
		t.Errorf("In correct computation between New York and New York: %v %v %v\n", int(miles), int64(km), ok)
	}
}
