type Test struct {
	On bool
	Ammo int
    Power int
}

func (t *Test) Shoot() bool {
    if(t.On && t.Ammo > 0){
        t.Ammo--
        
        return true
    }
    
	return false
}

func (t *Test) RideBike() bool {
    if(t.On && t.Power > 0){
        t.Power--
        
        return true
    }
    
	return false
}

func main() {
    testStruct := &Test{}