package service

import (
	"fmt"

	"Stay_watch/constant"
	"Stay_watch/model"

	"gorm.io/gorm"
)

type TagService struct{}

func (TagService) CreateTagMap(tagMap *model.TagMap) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Create(tagMap)
	if result.Error != nil {
		fmt.Printf("タグ登録処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

func (TagService) DeleteTagMap(userId int64) error {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return err
	}
	defer closer.Close()
	result := DbEngine.Where("user_id", userId).Unscoped().Delete(&model.TagMap{})
	if result.Error != nil {
		fmt.Printf("タグマップ削除処理失敗 %v", result.Error)
		return result.Error
	}
	return nil
}

func (TagService) GetTagMapIdsByUserId(userId int64) ([]int64, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tagMapIds := make([]int64, 0)
	result := DbEngine.Table("tag_maps").Where("user_id=?", userId).Select("id").Find(&tagMapIds)
	if result.Error != nil {
		fmt.Printf("タグID取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tagMapIds, nil
}

func (TagService) GetTagsByCommunityId(communityId int64) ([]model.Tag, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tags := make([]model.Tag, 0)

	// tagsテーブルからcommunity_idカラムがcommunityIdのnameの値をtagNamesに格納
	result := DbEngine.
		Where("community_id = ?", communityId).
		Where("EXISTS (?)", // tag_mapsによって1つもユーザにタグづけされていないタグは除外する
			DbEngine.
				Table("tag_maps").
				Select("1").                        // 存在するかしないかを知りたいだけで実際に何を返すかは重要じゃないのでとりあえず1
				Where("tag_maps.tag_id = tags.id"), // 外側クエリ（tags）の各行に対して、tag_id が一致する行があるかを確認
		).Find(&tags)
	if result.Error != nil {
		fmt.Printf("タグ一覧取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tags, nil
}

func (TagService) GetTagsByTagNames(tagNames []string, communityID int64) ([]model.Tag, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tags := make([]model.Tag, 0)

	// tagsテーブルから取得
	for _, tagName := range tagNames {
		tag := model.Tag{}
		var result *gorm.DB
		result = DbEngine.Where("name = ? AND (community_id = ? OR community_id = ?)", tagName, communityID, constant.PublicTagID).First(&tag)
		if result.Error != nil {
			// tagが見つからない場合新規作成する
			tag = model.Tag{
				Name:        tagName,
				CommunityId: communityID,
			}
			result = DbEngine.Create(&tag)
			if result.Error != nil {
				fmt.Println("タグ作成失敗")
				return nil, result.Error
			}
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (TagService) GetTagNamesByCommunityId(communityId int64) ([]string, error) {
	DbEngine := connect()
	closer, err := DbEngine.DB()
	if err != nil {
		return nil, err
	}
	defer closer.Close()
	tagNames := make([]string, 0)

	// tagsテーブルからcommunity_idカラムがcommunityIdのnameの値をtagNamesに格納
	result := DbEngine.Table("tags").Where("community_id=?", communityId).Select("name").Find(&tagNames)
	if result.Error != nil {
		fmt.Printf("タグ名一覧取得失敗 %v", result.Error)
		return nil, result.Error
	}

	return tagNames, nil
}
