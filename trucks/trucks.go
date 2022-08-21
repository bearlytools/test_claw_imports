// DO NOT EDIT
// This package is autogenerated and should not be modified except by the clawc compiler.

// Package trucks 
package trucks

import (
    "github.com/bearlytools/claw/languages/go/mapping"
    "github.com/bearlytools/claw/languages/go/reflect"
    "github.com/bearlytools/claw/languages/go/reflect/runtime"
    "github.com/bearlytools/claw/languages/go/structs"
    "github.com/bearlytools/claw/languages/go/field"
    
    "github.com/bearlytools/claw/testing/imports/vehicles/claw/manufacturers"
)

// SyntaxVersion is the major version of the Claw language that is being rendered.
const SyntaxVersion = 0


type Model uint8

// String implements fmt.Stringer.
func (x Model) String() string {
    return ModelByValue[uint8(x)]
}

// XXXEnumGroup will return the EnumGroup descriptor for this group of enumerators.
// This should only be used by the reflect package and is has no compatibility promises 
// like all XXX fields.
func (x Model) XXXEnumGroup() reflect.EnumGroup {
    return XXXEnumGroups.Get(0)
}

// XXXEnumGroup will return the EnumValueDescr descriptor for an enumerated value.
// This should only be used by the reflect package and is has no compatibility promises 
// like all XXX fields.
func (x Model) XXXEnumValueDescr() reflect.EnumValueDescr {
    return XXXEnumGroups.Get(0).ByValue(int(x))
}


const (
    ModelUnknown Model = 0
    F100 Model = 1
    Tundra Model = 2
    Cybertruck Model = 3
)

var ModelByName = map[string]Model{
    "Cybertruck": 3,
    "F100": 1,
    "ModelUnknown": 0,
    "Tundra": 2,
}

var ModelByValue = map[uint8 ]string{
    0: "ModelUnknown",
    1: "F100",
    2: "Tundra",
    3: "Cybertruck",
} 



type Truck struct {
   s *structs.Struct
}

// NewTruck creates a new instance of Truck.
func NewTruck() Truck {
    s := structs.New(0, mappingTruck)
    s.XXXSetNoZeroTypeCompression()
    return Truck{
        s: s,
    }
}

// XXXNewFrom creates a new Truck from our internal Struct representation.
// As with all things marked XXX*, this should not be used and has not compatibility
// guarantees.
//
// Deprecated: This is not actually deprecated, but it should not be used directly nor
// show up in any documentation.
func XXXNewFrom(s *structs.Struct) Truck {
    return Truck{s: s}
} 

func (x Truck) Manufacturer() manufacturers.Manufacturer {
    return manufacturers.Manufacturer(structs.MustGetNumber[uint8](x.s, 0))
}

func (x Truck) SetManufacturer(value manufacturers.Manufacturer) {
    structs.MustSetNumber(x.s, 0, uint8(value))
}  

func (x Truck) Model() Model {
    return Model(structs.MustGetNumber[uint8](x.s, 1))
}

func (x Truck) SetModel(value Model) {
    structs.MustSetNumber(x.s, 1, uint8(value))
} 


func (x Truck) Year() uint16 {
    return structs.MustGetNumber[uint16](x.s, 2)
}

func (x Truck) SetYear(value uint16) {
    structs.MustSetNumber(x.s, 2, value)
}  

// ClawStruct returns a reflection type representing the Struct.
func (x Truck) ClawStruct() reflect.Struct{
   return reflect.XXXNewStruct(x.s)
}

// XXXDescr returns the Struct's descriptor. This should only be used
// by the reflect package and is has no compatibility promises like all XXX fields.
func (x Truck) XXXDescr() reflect.StructDescr {
    return XXXPackageDescr.Structs()[0]
} 

// Everything below this line is internal details.
var mappingTruck = &mapping.Map{
    Name: "Truck",
    Pkg: "trucks",
    Path: "github.com/bearlytools/test_claw_imports/trucks",
    Fields: []*mapping.FieldDescr{
        {
            Name: "Manufacturer",
            Type: field.FTUint8,
            IsEnum: true,
            FieldNum: 0,
        },
        {
            Name: "Model",
            Type: field.FTUint8,
            IsEnum: true,
            FieldNum: 1,
        },
        {
            Name: "Year",
            Type: field.FTUint16,
            IsEnum: false,
            FieldNum: 2,
        },
    },
}
var _package = "trucks"

var XXXEnumGroups reflect.EnumGroups = reflect.XXXEnumGroupsImpl{
    List:   []reflect.EnumGroup{
        reflect.XXXEnumGroupImpl{
            GroupName: "Model",
            GroupLen: 4,
            EnumSize: 8,
            Descrs: []reflect.EnumValueDescr{
                reflect.XXXEnumValueDescrImpl{
                    EnumName: "ModelUnknown",
                    EnumNumber: 0,
                },
                reflect.XXXEnumValueDescrImpl{
                    EnumName: "F100",
                    EnumNumber: 1,
                },
                reflect.XXXEnumValueDescrImpl{
                    EnumName: "Tundra",
                    EnumNumber: 2,
                },
                reflect.XXXEnumValueDescrImpl{
                    EnumName: "Cybertruck",
                    EnumNumber: 3,
                },
            },
        },  
    },
    Lookup: map[string]reflect.EnumGroup{},
}

func init() {
    x := XXXEnumGroups.(reflect.XXXEnumGroupsImpl)
    for _, g := range x.List {
        x.Lookup[g.Name()] = g
    }
}  

var XXXPackageDescr reflect.PackageDescr = reflect.XXXPackageDescrImpl{
    Name: "trucks",
    Path: "github.com/bearlytools/test_claw_imports/trucks",
    ImportDescrs: []reflect.PackageDescr {
        manufacturers.XXXPackageDescr,  
    }, 
    EnumGroupsDescrs: XXXEnumGroups, 
    StructsDescrs: []reflect.StructDescr{
        reflect.XXXStructDescrImpl{
            Name: "Truck",
            Pkg: "trucks",
            Path: "github.com/bearlytools/test_claw_imports/trucks",
            FieldList: []reflect.FieldDescr{
                reflect.XXXFieldDescrImpl{
                    FD: mappingTruck.ByName("Manufacturer"),
                    EG: XXXEnumGroups.ByName("manufacturers.Manufacturer"),
                },
                reflect.XXXFieldDescrImpl{
                    FD: mappingTruck.ByName("Model"),
                    EG: XXXEnumGroups.ByName("Model"),
                },
                reflect.XXXFieldDescrImpl{
                    FD: mappingTruck.ByName("Year"),
                },
            },
        },  
    },  
}

func init() {
    runtime.RegisterPackage(XXXPackageDescr)
}