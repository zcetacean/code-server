package models

type BaseName struct {
  Name string `orm:"size(20)" json:"name" form:"name" valid:""`
}

type BaseId struct {
  Id string `orm:"size(50);pk" json:"id" form:"id" valid:""`
}

type BasePassword struct {
  Password string `orm:"size(50)" json:"password" form:"password" valid:""`
}

type BaseShouldScore struct {
  ShouldScore int `orm:"size(255)" json:"shouldScore" form:"shouldScore" valid:""`
}

type BaseApplyScore struct {
  ApplyScore int `orm:"size(255)" json:"applyScore" form:"applyScore" valid:""`
}

type BaseRealScore struct {
  RealScore int `orm:"size(255)" json:"realScore" form:"realScore" valid:""`
}

type BaseImg struct {
  Img string `orm:"size(255)" json:"img" form:"img" valid:""`
}

type BaseState struct {
  State string `orm:"size(10)" json:"state" form:"state" valid:""`
}

type BaseType struct {
  Type string `orm:"size(10)" json:"type" form:"type" valid:""`
}

type BasePermissionTreeName struct {
  PermissionTreeName string `orm:"size(20)" json:"permissionTreeName" form:"permissionTreeName" valid:""`
}

type BaseInitTreeId struct {
  InitTreeId string `orm:"size(50)" json:"initTreeId" form:"initTreeId" valid:""`
}

type BaseCurrentTreeId struct {
  CurrentTreeId string `orm:"size(50)" json:"currentTreeId" form:"currentTreeId" valid:""`
}

type BaseParentTreeId struct {
  ParentTreeId string `orm:"size(50)" json:"parentTreeId" form:"parentTreeId" valid:""`
}

type BaseRemark struct {
  Remark string `orm:"size(255)" json:"remark" form:"remark" valid:""`
}

type BaseTime struct {
  Time int `orm:"size(255)" json:"time" form:"time" valid:""`
}

type BaseInitTreeName struct {
  InitTreeName string `orm:"size(50)" json:"initTreeName" form:"initTreeName" valid:""`
}

type BaseMoralScore struct {
  MoralScore int `orm:"size(10)" json:"moralScore" form:"moralScore" valid:""`
}

type BaseAcademicScore struct {
  AcademicScore int `orm:"size(10)" json:"academicScore" form:"academicScore" valid:""`
}

type BasePhysicalScore struct {
  PhysicalScore int `orm:"size(10)" json:"physicalScore" form:"physicalScore" valid:""`
}
