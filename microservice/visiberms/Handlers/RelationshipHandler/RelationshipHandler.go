package visiberms_relationship_handler

import "github.com/vegh1010/golang.porfolio/microservice/visiberms/Features/Param"

type RelationshipHandler struct {
	visiberms_param.Param
}

func (self *RelationshipHandler) Init() {

}

func (self *RelationshipHandler) Close() {
	self.Param.Close()
}
