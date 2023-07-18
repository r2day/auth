package manage

import (
	"go.mongodb.org/mongo-driver/bson"
)

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}

// FindByPhone 通过手机号查找到账号信息
func (m *Model) FindByPhone(phone string) (*Model, error) {

	result := &Model{}
	filter := bson.D{{Key: "phone", Value: phone}}
	err := m.GetBy(result, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindByAccountId 通过手机号查找到账号信息
func (m *Model) FindByAccountId(accountID string) (*Model, error) {
	result := &Model{}
	filter := bson.D{{Key: "meta.account_id", Value: accountID}}
	err := m.GetBy(result, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateById 通过id更新数据库
// 直接使用mongo的id进行更新
// 这种情况一般用于先通过其他字段，例如phone 查找到记录
// 通过读取记录中的_id 进行更新
//func (m *Model) UpdateById(ctx context.Context) error {
//	coll := db.MDB.Collection(m.CollectionName())
//	// 更新数据库
//	m.Meta.UpdatedAt = rtime.FomratTimeAsReader(time.Now().Unix())
//	filter := bson.D{{Key: "_id", Value: m.ID}}
//	result, err := coll.UpdateOne(ctx, filter,
//		bson.D{{Key: "$set", Value: m}})
//	if err != nil {
//		return err
//	}
//	if result.MatchedCount < 1 {
//		log.WithField("id", m.ID).Warning("no matched record")
//		return nil
//	}
//	return nil
//}
