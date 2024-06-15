package generate

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pikachu0310/hackathon24spring02-data-server/internal/api"
	"github.com/pikachu0310/hackathon24spring02-data-server/openapi/models"
	"math/rand"
	"regexp"
)

var createItemText1 = `
僕は、AIを用いてアイテムを合成することができ、合成したアイテムを用いてオンラインで戦いあう、AI活用オンラインPvPゲームを作っています。
あなたはこのゲームの核である、アイテム生成の判断をするAIの役を担ってもらいます。

## まずこのゲームについての説明をします。

### 移動や操作について
プレイヤーはWASDでゆっくり移動することができます。そして、左クリックで弾を発射し、敵を攻撃することができます。ここまでは、diep.io のようなゲームです。しかし、diep.io にない特殊な操作が2種類あります。まず、右クリックをドラッグすることで、力をためて一気に加速することができます。そして、自分に最も近いオブジェクトに向かって糸を発射し、つかむことができます。

### ゲームの大きな特徴

#### 穴に落ちると死ぬ！
そして、さらにこのゲームの核でもある、とても面白い要素が二つあって、一つはマップ上の穴に落ちれば死ぬ、あるいは相手を穴に落とせば殺せるということです。
加速して相手をつかみ穴に落とす、といったようなオンライン対戦ならではのアクションが楽しめます。
このゲームは、かなり現実の物理演算に従って、正確に物理シミュレーションを行います。例えば、このゲームの移動は力を加えることで実現し、運動方程式に従い加速度が変化します。また、他のプレイヤーと衝突した際は運動量保存則に従います。

#### アイテムという要素があり、自機に合成して強化できる！
これがこのゲーム一番の凄くて面白い要素です。
プレイヤーは、マップに散らばっているアイテムをつかんで穴に落とすことで、アイテムを取得できます。アイテムを取得すると、インベントリに移動し、アイテムを自機に合成するか、アイテム同士を合成するか、捨てるかを選べます。自機に合成した場合、アイテムによって自分の様々なパラメーターが変化します。ただし、この時強くなるパラメーターがあれば、必ず弱くなるパラメーターもあります。つまり、メリットとデメリットが釣り合うようにアイテムの合成語のパラメーターを凄く考えなければなりません。
また、アイテム同士を合成する、ということもできます。弱いアイテムでも、合成によって使えるアイテムになるかもしれません。
プレイヤーは、各自がどのパラメーターを伸ばして、どのように戦うかを、このアイテムを合成したら自機がどのように強くなるかを予測しながら合成していく、という楽しみや知的好奇心を味わうことができます。
プレイヤーが事前に見ることができる情報は、アイテム名とそのアイテムのフレーバーテキストのみです。合成語のパラメーターは、AIであるあなたに決めてもらいます。

### パラメーター一覧
以下に、プレイヤーの持つパラメーター一覧を、効果: 変数名 (係数) のフォーマットで載せます。

- 自分の最大HP: maxHealth (1)
- 質量: mass (1)
- 反発係数: bounciness (0.4)
- 摩擦係数: friction (0.3)
- 大きさ: size (1)
- HP回復速度: hpRegenSpeed (1)
- 防御力: defence (1)
- WASDで与える力の大きさ: movementSpeed (1)
- 突き操作のクールダウン時間: dashCooldownTime (1)
- 突き操作の矢印がたまる速度: dashArrowFillRate (1)
- 突き操作の矢印の長さ: dashArrowMaxLength (1)
- 突き操作の加速の時間の長さ: dashAccelerationDuration (1)
- 突き操作で与えられる力の最大の大きさ: dashMaxForce (1)
- 弾の攻撃力: bulletAttack  (1)
- 同時に発車する弾の数: bulletNumber (1)
- 弾の散乱角度: bulletAngle (1)
- 弾の速度: bulletSpeed (1)
- 弾が消滅するまでの時間の長さ: bulletAliveTime (1)
- 弾の連射速度: bulletInterval (1)
- 弾の大きさ: bulletSize (1)

## AIであるあなたにやってもらいたい事
アイテムを1個生成してもらいます。アイテムは、名前とフレーバーテキストの二つだけの情報を持っています。
アイテムを出力するときは、必ず以下のフォーマットで出力してください。

name: 弾性力の強いゴム
description: とても弾性力の強いゴムだ。高い所から投げると、スーパーボールのように跳ね回る。これで数時間は遊べそうだ。

ただし、アイテムを生成するときに注意点があります。
まず、これはゲームなので、合成意欲を掻き立てるようなアイテムを作ってください。また、フレーバーテキストは、真面目な文章と、ユニークで面白い文章を混ぜてください。
また、メリットが強いアイテムには、相応のデメリットが必要であることを意識してください。
アイテムはAIによる生成だけであり、合成したらどうなるか分からないため、アイテム合成をさせた結果前より弱くなったり、またはとても強くなったりなど、どんなアイテムが出来るか分からないという楽しみがプレイヤーにあります。

それぞれの情報に対して、パラメータについてより詳細に教えます。
- Name : アイテムの名前です。合成意欲を掻き立てるような名前のアイテムで、かつ簡潔で短い名前が望ましいです。
- Description : アイテムの説明文です。こちらは、結構長めで、真面目な文章と、ユニークで面白い文章を混ぜて楽しい感じにしてください。これはアイテム合成の際のパラメーターを算出する際に、かなり参考にします。なので、それを考慮したうえで、そのアイテムの特徴をしっかりと書いてください。

### 出力のフォーマット
アイテムを出力する際は、以下のフォーマットで出力してください。

name: 弾性力の強いゴム
description: とても弾性力の強いゴムだ。高い所から投げると、スーパーボールのように跳ね回る。これで数時間は遊べそうだ。

必ずアイテムを出力するときは上のようなフォーマットで、アイテムの情報以外何も書かずにアイテムの情報だけを出力してください。

### アイテム生成の詳細
アイテム生成の注意点として、そのアイテム単体だとあまり役に立たなかったりとても弱かったりするが、合成意欲を掻き立てるようなアイテムを作ってほしいです。例えば、水や氷、草や炎といった属性っぽさがありそうなアイテム名を付けることが出来れば、プレイヤーは属性っぽさからヒントを得て面白い合成を思いつくかもしれません。例えば、氷っぽいアイテム「凍った土」と水っぽいアイテム「水鉄砲」を組み合わせれば、水を凍らせられて強いアイテムができるのではないかとか考えるかもしれません。また、属性に限らず、形容詞を付けてあげるといいかもしれません。つまり、"弱いけど組み合わせたら強くなるかも"なアイテム名が面白いと思います。
アイテムは、自機に合成して自機を強化するための素材です。先ほど列挙したパラメーターに活かせそうなアイテムが望ましいです。

改めて、以上のようなことを踏まえて、アイテムを1個考えて生成して、決まったフォーマットに従って出力してください。
`

func CreateItem() (item *models.Item, err error) {
	messages := api.CreateNewMessages()
	api.AddMessageAsUser(messages, createItemText1)
	responseText, reason, err := api.RequestGPTAndGetResponseText(messages)
	fmt.Println("****AI OUTPUT****\n" + responseText)
	if err != nil || reason == api.ErrorHappen {
		fmt.Println("GPT ERROR:" + err.Error())
		return nil, err
	} else if reason == api.Length {
		fmt.Println("GPT ERROR: LENGTH")
		return nil, err
	}

	return parseItem(responseText)
}

//func parseItems(s string) ([]*Item, error) {
//	itemTextsInput := strings.Split(s, "```")
//	var itemTexts []string
//	for _, itemText := range itemTextsInput {
//		if len(strings.Split(itemText, "\n")) >= 12 {
//			return make([]*Item, 0), fmt.Errorf("Invalid input format3" + itemText)
//		}
//		if len(strings.Split(itemText, "\n")) >= 9 {
//			itemTexts = append(itemTexts, itemText)
//		}
//	}
//	items := make([]*Item, 0)
//	for _, itemText := range itemTexts {
//		item, err := parseItem(itemText)
//		if err != nil {
//			return make([]*Item, 0), err
//		}
//		items = append(items, item)
//	}
//	return items, nil
//}

func parseItem(text string) (*models.Item, error) {
	// 正規表現パターン
	re := regexp.MustCompile(`(?s)name:\s*(.*?)\s*description:\s*(.*?)\s*$`)

	// マッチを検索
	matches := re.FindStringSubmatch(text)
	if matches == nil {
		fmt.Println("No match found")
		return nil, errors.New("ParseItem: Not Match Found")
	}

	// 抽出された内容を変数に代入
	name := matches[1]
	description := matches[2]

	createdItem := &models.Item{
		Attribute:   rand.Intn(6), //TODO
		Description: description,
		Id:          uuid.New(),
		Name:        name,
		Rarity:      rand.Intn(3), //TODO
	}

	return createdItem, nil
}

//func GptGenerateItem() ([]*Item, error) {
//	requestContent = MakeItemMessages
//	responseText, reason, err := api.RequestGPTAndGetResponseText(MakeItemMessages)
//	fmt.Println("****AI OUTPUT****\n" + responseText)
//	if err != nil || reason == api.ErrorHappen {
//		fmt.Println("GPT ERROR:" + err.Error())
//		return nil, err
//	} else if reason == api.Length {
//		fmt.Println("GPT ERROR: LENGTH")
//		return nil, err
//	}
//
//	return parseItems(responseText)
//}
//
//func GptCombineItem(items []*Item) ([]*Item, error) {
//	requestContent = CombineItemMessages
//	CombineItemMessageTemp := CombineItemMessage3.Content
//	for i, item := range items {
//		CombineItemMessageTemp += fmt.Sprintf("アイテム%d```\n%s\n%s\n%d\n%d\n%d\n%d\n%d\n```\n", i, item.Name, item.Category, item.MaxHp, item.InstantHeal, item.SustainedHeal, item.Attack, item.Defense)
//	}
//	res, err := Gpt(CombineItemMessageTemp, func(s string) {})
//	fmt.Println("****AI OUTPUT****\n" + res.Text())
//	if err != nil {
//		return nil, err
//	}
//	if len(res.Text()) >= 7 && res.Text()[:7] == "error:" {
//		return nil, err
//	}
//	return parseItems(res.Text())
//}
//
//func SendOrEditError(SendOrEdit func(string), err error) {
//	SendOrEdit(fmt.Sprintf("error: %v", err))
//}
//
//func Gpt(content string, SendOrEdit func(string)) (api.OpenaiResponse, error) {
//	addRequestContent("user", content)
//	res, err := api.RequestOpenaiApiByMessages(requestContent)
//	if err != nil {
//		SendOrEditError(SendOrEdit, err)
//		return res, err
//	}
//	res, err = GptDeleteLogsAndRetry(res, SendOrEdit)
//	if err != nil {
//		SendOrEditError(SendOrEdit, err)
//		return res, err
//	}
//	addRequestContent("assistant", res.Text())
//	responses = append(responses, res)
//	SendOrEdit(res.Text())
//	return res, err
//}
//
//func GptDeleteLogsAndRetry(res api.OpenaiResponse, SendOrEdit func(string)) (api.OpenaiResponse, error) {
//	var err error
//	for i := 0; res.OverTokenCheck() && i <= 4; i++ {
//		SendOrEdit("Clearing old history and retrying.[" + fmt.Sprintf("%d", i+1) + "] :thinking:")
//		if len(requestContent) >= 5 {
//			requestContent = requestContent[4:]
//			requestContent = append([]api.Message{firstMessage}, requestContent[4:]...)
//		} else if len(requestContent) >= 2 {
//			requestContent = append([]api.Message{firstMessage}, requestContent[1:]...)
//		} else if len(requestContent) >= 1 {
//			requestContent = []api.Message{firstMessage}
//		}
//		res, err = api.RequestOpenaiApiByMessages(requestContent)
//		if err != nil {
//			SendOrEditError(SendOrEdit, err)
//		}
//	}
//	return res, err
//}
//
//func GptReset(SendOrEdit func(string)) (api.OpenaiResponse, error) {
//	resetRequestContent()
//	// resetResponses()
//	// res, err := api.RequestOpenaiApiByStringOneTime(ResetMessage)
//	// if err != nil {
//	//	SendOrEditError(SendOrEdit, err)
//	//	return res, err
//	// }
//	// SendOrEdit(res.Text())
//	// return res, err
//	return api.OpenaiResponse{}, nil
//}
//
//func Sum(arr []float32) float32 {
//	var res float32 = 0
//	for i := 0; i < len(arr); i++ {
//		res += arr[i]
//	}
//	return res
//}
