package models

type Organize struct {
  BaseId
  BaseName
  BaseState
  BaseParentTreeId
}
type Item struct {
  BaseId
  BaseName
  BaseShouldScore
  BaseInitTreeId
  BaseType
}
type Verify struct {
  BaseId
  BaseState
  BaseImg
  BaseRealScore
  BaseApplyScore
  BaseRemark
  BaseType
  BaseTime
  BaseCurrentTreeId
  Student *Student `orm:"rel(fk)"`
  Item *Item `orm:"rel(fk)"`
}
type Admin struct {
  BaseId
  BaseName
  BasePassword
  BaseState
  BasePermissionTreeName
}
type Record struct {
  BaseId
  BaseTime
  BaseRemark
  Admin *Admin `orm:"rel(fk)"`
}
type Student struct {
  BaseId
  BasePassword
  BaseName
  BaseInitTreeName
  BaseMoralScore
  BaseAcademicScore
  BasePhysicalScore
}
