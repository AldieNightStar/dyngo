package dyngo

const (
	TYPE_NIL = iota
	TYPE_INT
	TYPE_FLOAT
	TYPE_STR
	TYPE_RUNE
	TYPE_BOOL
	TYPE_ARR
	TYPE_MAP
	TYPE_FUNC
)

type VarFunc func(v ...*Var) *Var

type Var struct {
	Type    int
	Integer int
	Float   float64
	String  string
	Rune    rune
	Boolean bool
	Array   []*Var
	Map     map[string]*Var
	Func    VarFunc
}

func VarOfNil() *Var {
	return &Var{Type: TYPE_NIL}
}

func VarOfInt(n int) *Var {
	return &Var{Type: TYPE_INT, Integer: n}
}

func VarOfFloat(f float64) *Var {
	return &Var{Type: TYPE_FLOAT, Float: f}
}

func VarOfString(s string) *Var {
	return &Var{Type: TYPE_STR, String: s}
}

func VarOfRune(r rune) *Var {
	return &Var{Type: TYPE_RUNE, Rune: r}
}

func VarOfBool(b bool) *Var {
	return &Var{Type: TYPE_BOOL, Boolean: b}
}

func VarOfArr(a []*Var) *Var {
	return &Var{Type: TYPE_ARR, Array: a}
}

func VarOfMap(m map[string]*Var) *Var {
	return &Var{Type: TYPE_MAP, Map: m}
}

func NewVarArray(v ...*Var) *Var {
	return VarOfArr(v)
}

// =============================
// Methods (GETTERS)
// =============================

func (v *Var) IsTrue() bool {
	return v.Type == TYPE_BOOL && v.Boolean
}

func (v *Var) GetInt(def int) int {
	if v.Type == TYPE_INT {
		return v.Integer
	}
	return def
}

func (v *Var) GetFloat(def float64) float64 {
	if v.Type == TYPE_FLOAT {
		return v.Float
	}
	return def
}

func (v *Var) GetString(def string) string {
	if v.Type == TYPE_STR {
		return v.String
	}
	return def
}

func (v *Var) GetRune(def rune) rune {
	if v.Type == TYPE_RUNE {
		return v.Rune
	}
	return def
}

func (v *Var) GetBool(def bool) bool {
	if v.Type == TYPE_BOOL {
		return v.Boolean
	}
	return def
}

func (v *Var) GetArray(def []*Var) []*Var {
	if v.Type == TYPE_ARR {
		return v.Array
	}
	return def
}

func (v *Var) GetMap(def map[string]*Var) map[string]*Var {
	if v.Type == TYPE_MAP {
		return v.Map
	}
	return def
}

func (v *Var) GetFunc(def VarFunc) VarFunc {
	if v.Type == TYPE_MAP {
		return v.Func
	}
	return def
}

// =============================
// Methods (GETTERS)
// =============================

func (v *Var) SetInt(n int) {
	v.Type = TYPE_INT
	v.Integer = n
}

func (v *Var) SetFloat(n float64) {
	v.Type = TYPE_FLOAT
	v.Float = n
}

func (v *Var) SetString(s string) {
	v.Type = TYPE_STR
	v.String = s
}

func (v *Var) SetRune(r rune) {
	v.Type = TYPE_RUNE
	v.Rune = r
}

func (v *Var) SetBool(b bool) {
	v.Type = TYPE_BOOL
	v.Boolean = b
}

func (v *Var) SetArray(a []*Var) {
	v.Type = TYPE_ARR
	v.Array = a
}

func (v *Var) SetMap(m map[string]*Var) {
	v.Type = TYPE_MAP
	v.Map = m
}

func (v *Var) SetFunc(f VarFunc) {
	v.Type = TYPE_FUNC
	v.Func = f
}

// =============================
// Methods (Map)
// =============================

func (v *Var) Len() int {
	if v.Type == TYPE_ARR {
		return len(v.Array)
	} else if v.Type == TYPE_MAP {
		return len(v.MapKeys())
	} else if v.Type == TYPE_STR {
		return len(v.String)
	}
	return 0
}

func (v *Var) MapGet(k string) *Var {
	mp := v.GetMap(nil)
	if mp == nil {
		return VarOfNil()
	}
	e, ok := mp[k]
	if !ok {
		return VarOfNil()
	}
	return e
}

func (v *Var) MapSet(k string, val *Var) bool {
	mp := v.GetMap(nil)
	if mp == nil {
		return false
	}
	mp[k] = val
	return true
}

func (v *Var) MapKeys() []string {
	arr := make([]string, 0, 8)
	mp := v.GetMap(nil)
	if mp == nil {
		return []string{}
	}
	for k := range mp {
		arr = append(arr, k)
	}
	return arr
}

// =============================
// Methods (Array)
// =============================

func (v *Var) ArrGet(id int, def *Var) *Var {
	arr := v.GetArray(nil)
	if arr == nil {
		return def
	}
	if id < 0 || id >= len(arr) {
		return def
	}
	return arr[id]
}

func (v *Var) ArrSet(id int, val *Var) bool {
	arr := v.GetArray(nil)
	if arr == nil {
		return false
	}
	if id < 0 || id >= len(arr) {
		return false
	}
	arr[id] = val
	return true
}

func (v *Var) ArrAdd(val *Var) bool {
	arr := v.GetArray(nil)
	if arr == nil {
		return false
	}
	arr = append(arr, val)
	v.SetArray(arr)
	return true
}

func (v *Var) ArrClear() bool {
	if v.Type == TYPE_ARR {
		v.Array = make([]*Var, 0, 8)
		return true
	}
	return false
}

// =============================
// Methods (Func)
// =============================

func (v *Var) Call(args ...*Var) *Var {
	if v.Type == TYPE_FUNC {
		return v.Func(args...)
	}
	return VarOfNil()
}

// =============================
// Methods (Simple Math)
// =============================

func (v *Var) Add(n *Var) *Var {
	if v.Type == TYPE_INT {
		if n.Type == TYPE_INT {
			return VarOfInt(v.Integer + n.Integer)
		} else if n.Type == TYPE_FLOAT {
			return VarOfFloat(float64(v.Integer) + n.Float)
		}
	} else if v.Type == TYPE_FLOAT {
		if n.Type == TYPE_FLOAT {
			return VarOfFloat(v.Float + n.Float)
		} else if n.Type == TYPE_INT {
			return VarOfFloat(v.Float + float64(n.Integer))
		}
	}
	return v
}

func (v *Var) Sub(n *Var) *Var {
	if v.Type == TYPE_INT {
		if n.Type == TYPE_INT {
			return VarOfInt(v.Integer - n.Integer)
		} else if n.Type == TYPE_FLOAT {
			return VarOfFloat(float64(v.Integer) - n.Float)
		}
	} else if v.Type == TYPE_FLOAT {
		if n.Type == TYPE_FLOAT {
			return VarOfFloat(v.Float - n.Float)
		} else if n.Type == TYPE_INT {
			return VarOfFloat(v.Float - float64(n.Integer))
		}
	}
	return v
}

func (v *Var) Div(n *Var) *Var {
	if v.Type == TYPE_INT {
		if n.Type == TYPE_INT {
			return VarOfInt(v.Integer / n.Integer)
		} else if n.Type == TYPE_FLOAT {
			return VarOfFloat(float64(v.Integer) / n.Float)
		}
	} else if v.Type == TYPE_FLOAT {
		if n.Type == TYPE_FLOAT {
			return VarOfFloat(v.Float / n.Float)
		} else if n.Type == TYPE_INT {
			return VarOfFloat(v.Float / float64(n.Integer))
		}
	}
	return v
}

func (v *Var) Mul(n *Var) *Var {
	if v.Type == TYPE_INT {
		if n.Type == TYPE_INT {
			return VarOfInt(v.Integer * n.Integer)
		} else if n.Type == TYPE_FLOAT {
			return VarOfFloat(float64(v.Integer) * n.Float)
		}
	} else if v.Type == TYPE_FLOAT {
		if n.Type == TYPE_FLOAT {
			return VarOfFloat(v.Float * n.Float)
		} else if n.Type == TYPE_INT {
			return VarOfFloat(v.Float * float64(n.Integer))
		}
	}
	return v
}

func (v *Var) Mod(n *Var) *Var {
	if v.Type == TYPE_INT {
		if n.Type == TYPE_INT {
			return VarOfInt(v.Integer % n.Integer)
		} else if n.Type == TYPE_FLOAT {
			return VarOfInt(v.Integer % int(n.Float))
		}
	} else if v.Type == TYPE_FLOAT {
		if n.Type == TYPE_FLOAT {
			return VarOfInt(int(v.Float) % int(n.Float))
		} else if n.Type == TYPE_INT {
			return VarOfInt(int(v.Float) % n.Integer)
		}
	}
	return v
}

// =============================
// Methods (Functional)
// =============================

func (v *Var) ForEach(arg *Var) {
	f := arg.GetFunc(nil)
	if f == nil {
		return
	}
	if v.Type == TYPE_ARR {
		arr := v.GetArray(nil)
		if arr == nil {
			return
		}
		for id := range arr {
			f(VarOfInt(id), arr[id])
		}
	} else if v.Type == TYPE_MAP {
		mp := v.GetMap(nil)
		if mp == nil {
			return
		}
		for k := range mp {
			f(VarOfString(k), mp[k])
		}
	} else if v.Type == TYPE_STR {
		s := v.GetString("")
		if s == "" {
			return
		}
		for i := 0; i < len(s); i++ {
			f(VarOfInt(i), VarOfRune(rune(s[i])))
		}
	}
}
