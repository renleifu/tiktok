package dao

import (
	"tiktok/main/common"
	. "tiktok/main/config"
)

func deleteRelationByUerA(from_user_id int64, to_user_id int64) {
	relationship_follow := common.RelationshipFollow{}
	DB.Where("follow_id = ? and follower_id = ? ", from_user_id, to_user_id).Delete(&relationship_follow)

}
