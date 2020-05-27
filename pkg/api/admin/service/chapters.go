package service

import (
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"lime/pkg/api/admin/dao"
	"lime/pkg/api/admin/dto"
	"lime/pkg/api/admin/model"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

var ChaptersDao = dao.ChaptersDao{}

// Service
type ChaptersService struct {
}

// InfoOfId
func (bs ChaptersService) InfoOfId(dto dto.GeneralGetDto) model.Chapters {
	return ChaptersDao.Get(dto.Id)
}

func (bs ChaptersService) GetAll() []model.Chapters {
	return ChaptersDao.GetAll()
}

// List
func (bs ChaptersService) List(dto dto.GeneralListDto) ([]model.Chapters, int64) {
	return ChaptersDao.List(dto)
}

// Create
func (bs ChaptersService) Create(dto dto.ChaptersCreateDto) (model.Chapters, error) {
	Model := model.Chapters{
		Novel_id:   dto.Novel_id,
		Chapter_no: dto.Chapter_no,
		Title:      dto.Title,
		Desc:       dto.Desc,
		Link:       dto.Link,
		Is_vip:     dto.Is_vip,
		Source:     dto.Source,
		Views:      dto.Views,
		Text_num:   dto.Text_num,
		Status:     dto.Status,
		Try_views:  dto.Try_views,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	c := ChaptersDao.Create(&Model)
	if c.Error != nil {
		log.Error(c.Error.Error())
	}
	return Model, nil
}

// Update
func (bs ChaptersService) Update(dto dto.ChaptersEditDto) int64 {
	Model := model.Chapters{
		Id:         dto.Id,
		Novel_id:   dto.Novel_id,
		Chapter_no: dto.Chapter_no,
		Title:      dto.Title,
		Desc:       dto.Desc,
		Link:       dto.Link,
		Is_vip:     dto.Is_vip,
		Source:     dto.Source,
		Views:      dto.Views,
		Text_num:   dto.Text_num,
		Status:     dto.Status,
		Try_views:  dto.Try_views,
		UpdatedAt:  time.Now(),
	}
	c := ChaptersDao.Update(&Model, map[string]interface{}{
		"novel_id":   dto.Novel_id,
		"chapter_no": dto.Chapter_no,
		"title":      dto.Title,
		"desc":       dto.Desc,
		"link":       dto.Link,
		"is_vip":     dto.Is_vip,
		"source":     dto.Source,
		"views":      dto.Views,
		"text_num":   dto.Text_num,
		"status":     dto.Status,
		"try_views":  dto.Try_views,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

// Update
func (bs ChaptersService) UpdateStatus(dto dto.ChaptersUpdateStatusDto) int64 {
	Model := model.Chapters{
		Id: dto.Id,
	}
	c := ChaptersDao.Update(&Model, map[string]interface{}{
		"status":     dto.Status,
		"updated_at": time.Now(),
	})
	return c.RowsAffected
}

//Delete
func (bs ChaptersService) Delete(dto dto.GeneralDelDto) int64 {
	Model := model.Chapters{
		Id: dto.Id,
	}
	c := ChaptersDao.Delete(&Model)
	return c.RowsAffected
}

//上传封面
func (bs ChaptersService) UploadCover(file multipart.File, filename string) (filepath string, err error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return "", err
	}
	file.Close()
	// 大图
	m := resize.Resize(225, 300, img, resize.Lanczos3)
	filename = strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	out, err := os.Create("./data/images/" + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)

	//缩略图
	dst := "./data/images/" + strings.Replace(filename, ".", "_small.", 1)
	mSmall := resize.Resize(75, 100, img, resize.Lanczos3)
	outSmall, errSmall := os.Create(dst)
	if errSmall != nil {
		return "", errSmall
	}
	defer outSmall.Close()
	jpeg.Encode(outSmall, mSmall, nil)

	filepath = "/upload/images/" + filename
	return filepath, nil
}

func (bs ChaptersService) GetChapterInfoById(bdto dto.GeneralGetDto) dto.ChapterInfoDto {
	return dto.ChapterInfoDto{
		Book: dto.ChapterBookInfo{
			BookId:     1,
			Name:       "道是无情",
			IsFinished: 1,
		},
		ChapterTitle: "第1章 我们是兄妹！",
		Content:      "\t    明天就是男朋友秦牧扬的生日，我漂洋过海来给他过生日，本想给他一个惊喜，却未曾想惊喜倒没有，倒是给自己一个大大的惊吓。n\n\tn\n\t    下了飞机，坐上一辆出租车，人若倒霉喝凉水都塞牙，出租车走半道翻车了，擦伤了手臂，被警察送往医院的时候，我反应过来正想给秦牧扬打电话时，迎面走来一个男人搀扶着一个孕妇。n\n\tn\n\t    大着肚子挺的很高一看就像是至少有五个多月的身子。n\n\tn\n\t    我现在关注的不是那个孕妇。n\n\tn\n\t    而是孕妇身边的那个男人，那个男人不是别人，正是我的男朋友秦牧扬。n\n\tn\n\t    秦牧扬对那个孕妇小心翼翼的模样，女人的第六感告诉我秦牧扬和那个女人，关系绝不简单。n\n\tn\n\t    我走到秦牧扬的面前，尽量平复自己的心情，声音还算正常“你这是在干嘛！”n\n\tn\n\t    秦牧扬一脸意外的模样，他很意外我怎么会出现在这里。n\n\tn\n\t    震惊失措躲闪，再看看他身边的这个女人，挑衅的看着我，我的心里隐隐的像是明白了什么。n\n\tn\n\t    “秦牧扬，我在问你话呢？”n\n\tn\n\t    我的口气有些咄咄逼人。n\n\tn\n\t    秦牧扬没有立马就回答我的话，他看了看身边的女人一眼，就拉着我的手走，走了大概几十米米远我甩开他的的手，面部狰狞像个疯子一样愤怒的吼道“那个女人跟你是什么关系，他肚子里的孩子是你的吗？”n\n\tn\n\t    他眉头紧皱，眼神里闪烁着愧疚。n\n\tn\n\t    我和他相识18年，他的一个眼神，我就知道他心里的想法是什么。n\n\tn\n\t    “木子，对不起！”n\n\tn\n\t    秦牧扬跟我说了五个字，五个字对于我来说却像是天打雷劈一般。n\n\tn\n\t    他觉得我伤的不够深似的，顿了几秒又道“木子，我们分手吧！”n\n\tn\n\t    我踉跄的后退了几步，似乎不敢相信自己耳朵里听见的话。n\n\tn\n\t    “啪！”我重重的一巴掌甩在他的脸上，他没有丝毫躲闪的意思，接了这巴掌。n\n\tn\n\t    “你……你……欺人太甚！”我再也受不了眼泪像是决了堤的洪水似的一涌而出。n\n\tn\n\t    我们之间的恋爱关系，建立时间也不过三个月的时间而已，而那个女人的肚子却有五个多月大，也就是说他跟别的女人都有孩子了，三个月前他却跟我表白，让我做他的女朋友。n\n\tn\n\t    原来我就像是一个傻子一样被他耍的团团转。n\n\tn\n\t    他是我最爱是人啊，他是这十八年来里，对我最好的人啊，可如今却欺骗了我，玩弄了我。我此时此刻已经无法去形容我心中的愤怒，手不受控制的抬起重重的一巴掌再次的甩在他脸上。n\n\tn\n\t    同样，他依然没有躲开。n\n\tn\n\t    倒是那个女人过来，上来就狠狠的给了我一巴掌。n\n\tn\n\t    “我和他是在你之前，你有什么资格打他，要是论三儿，那个小三儿也是你。”n\n\tn\n\t    那个女人说我是她和秦牧扬之间的小三儿吗？n\n\tn\n\t    我捂着自己的脸，秦牧扬也震惊这个女人打了我。n\n\tn\n\t    “魏冉你在干什么，我们之间的事情不需要你来管，你先回去。”他吼着那个女人。n\n\tn\n\t    那个叫魏冉的女人见秦牧扬对她发了火，也不打算跟我纠缠什么就要走，我拽着她的手肯定不让她走“事情没有说清楚，谁都别想走。”n\n\tn\n\t    “好啊！那我不介意跟李小姐说个清楚道个明白，让你清楚谁才是小三儿。”n\n\tn\n\t    魏冉也不是什么善类。n\n\tn\n\t    “够了，魏冉你闭嘴。”秦牧扬再次朝那个女人发火。n\n\tn\n\t    “木子，对不起，这事儿我会慢慢跟你解释清楚，你现在住哪家酒店，我送你回去。”n\n\tn\n\t    秦牧扬脸上疲倦的神色，他累了，显然是两个女人折腾的他身心疲倦。n\n\tn\n\t    “我要你的解释，现在！”n\n\tn\n\t    此刻，我愤怒得像是失去了理智，像是一个动物一般只想着发泄，我一把甩开魏冉的手，双手双脚就要往他身上发泄去。n\n\tn\n\t    他不躲，那个女人上来护着他，我愤怒到没有丝毫理智，一把推开她。n\n\tn\n\t    “啊！”n\n\tn\n\t    女人一声儿尖叫。n\n\tn\n\t    “魏冉魏冉，你怎么了？”n\n\tn\n\t    “牧扬，我肚子好疼好疼，我可能要流产了，牧扬我们的孩子。”n\n\tn\n\t    他拽着我的手，就将我甩在地上跪着。n\n\tn\n\t    我看着他着急的神色，抱起那个女人，他手上都是血，他疯狂的往医院里冲。n\n\tn\n\t    我的膝盖被碎石子扎破了，在疼也不及心疼。n\n\tn\n\t    那个女人趟过的地方大片的血，我反应过来，看着自己的双手才意识到自己刚才都做了什么。n\n\tn\n\t    医院的走廊里，秦牧扬依靠在墙壁上，一根又一根的香烟抽着，浑身散发着颓废的气息。n\n\tn\n\t    他在国外留学这几年，每年我都要去看他几次，竟然不知道他什么时候染上了抽烟的恶习。n\n\tn\n\t    通过他简单的叙述我知道了他和那个魏冉是怎么回事，他和魏冉之前参加过一次聚会，不小心发生了一夜情，魏冉怀孕了却没有告诉他，等肚子都三四个月了才跟他说。n\n\tn\n\t    也就是说，他跟我表白的时候，并不知道那个女人怀了他的孩子。n\n\tn\n\t    他没有欺骗我，也没有玩弄我。n\n\tn\n\t    “木子，你回去吧！是我对不住你。”n\n\tn\n\t    他抽完又一根烟，抬眸对我说。n\n\tn\n\t    他还是要跟我分手，可是我不同意，我怎么会同意呢，我是那么的爱他啊！n\n\tn\n\t    我爱他已经爱到骨子里了，我不能没有他。n\n\tn\n\t    我听见自己的声音都在颤抖“你……可以让她把这个孩子生下来，也可以和她一起抚养这个孩子，我不介意的，我也会帮着你一起照顾，甚至我可以不生自己的孩子，只要你娶我别娶她！好吗？”我蹲在他面前，将脑袋埋在他的腿上，卑微的祈求。n\n\tn\n\t    我从来没有这般卑微过。n\n\tn\n\t    他知道，我是个怎样的性子。n\n\tn\n\t    这样的卑微我从未有过。n\n\tn\n\t    秦牧扬抬起我的头，震惊的看着我，他的手指轻柔的捻去我眼角的泪“木子，对不起，我不值得你这样，不值得！”n\n\tn\n\t    “值得，值得的，你可知道我有多爱你，你比谁都知道，我等了你这么多年，终于等到你的表白，我不想放弃你，求你了，不要，别不要我好吗！”n\n\tn\n\t    我也也没有了刚才打人的那种张牙舞爪的气势，现在有的只是脆弱。n\n\tn\n\t    我怕这个男人不要我。n\n\tn\n\t    他轻轻将我推开。n\n\tn\n\t    我不甘心，从后面抱住他的腰，将脸埋在他的后背不停的蹭着“我不介意，我真的不介意！”n\n\tn\n\t    可是他还是掰开了我的手“就算没有魏然肚子里的孩子，我们也走不到最后，别忘了我们是兄妹！”n\n\t",
		IsPreview:    0,
		LastChapter:  0,
		NextChapter:  "62759",
		ChapterId:    62758,
		UpdateTime:   "1536635791",
	}
}

// List
func (bs ChaptersService) FrontList(gdto dto.GeneralListDto) dto.ChaptersFrontListDto {
	return dto.ChaptersFrontListDto{
		BookId:       0,
		Name:         "道是无情",
		Cover:        "http://beiwo-new.oss-cn-beijing.aliyuncs.com/cover/222/3ddb80209a5db297473192bab469a1ef.jpeg?x-oss-process=image%2Fresize%2Cw_143%2Ch_200%2Cm_lfit",
		Description:  "十年前，我拿剪刀戳伤了他的眉心，我被关了半个月的少改所。十年后，他指着我说我下贱，勾引了他的弟弟。我和他相看两厌，他厌恶我，我痛恨他。本以为就这样一辈子在无交集。一天，他却装醉欺辱了我...我恨他，致死！",
		TotalChapter: 160,
		ChapterList:  []dto.ChaptersFrontLists{
			{
				BookId:1,
				ChapterId: 1,
				ChapterTitle:"",
				Words: 100,
				IsVip: 1,
				CanRead: 1,
				IsPreview:0,
				UpdateTime: "1536635791",
				DisplayOrder: 0,
				LastChapter: 0,
				NextChapter: 2,
				Tag: dto.BookInfoTag{
					Tab:   "免费",
					Color: "#778899",
				},
			},
		},
	}
}

