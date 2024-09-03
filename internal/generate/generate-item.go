package generate

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pikachu0310/hackathon24spring02-data-server/internal/api"
	"github.com/pikachu0310/hackathon24spring02-data-server/openapi/models"
	"math/rand"
	"regexp"
	"strconv"
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

var combineText = `
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

#### アイテムという要素があり、自機に合成したりアイテム同士を合成して強化できる！
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
アイテムを2個与えるので、あなたに合成してもらいます。アイテムは、名前とフレーバーテキストの二つだけの情報を持っています。
アイテムを出力するときは、必ず以下のフォーマットで出力してください。

name: 弾性力の強いゴム
description: とても弾性力の強いゴムだ。高い所から投げると、スーパーボールのように跳ね回る。これで数時間は遊べそうだ。

ただし、アイテムを合成するときに注意点があります。
まず、これはゲームなので、合成意欲を掻き立てるようなアイテムを作ってください。また、フレーバーテキストは、真面目な文章と、ユニークで面白い文章を混ぜてください。
また、メリットが強いアイテムには、相応のデメリットが必要であることを意識してください。
みんな合成したらどうなるか分からないため、アイテム合成をさせた結果前より弱くなったり、またはとても強くなったりなど、どんなアイテムが出来るか分からないという楽しみがプレイヤーにあります。

それぞれの情報に対して、パラメータについてより詳細に教えます。
- Name : アイテムの名前です。合成意欲を掻き立てるような名前のアイテムで、かつ簡潔で短い名前が望ましいです。
- Description : アイテムの説明文です。こちらは、結構長めで、真面目な文章と、ユニークで面白い文章を混ぜて楽しい感じにしてください。これはアイテム合成の際のパラメーターを算出する際に、かなり参考にします。なので、それを考慮したうえで、そのアイテムの特徴をしっかりと書いてください。

### 出力のフォーマット
アイテムを出力する際は、以下のフォーマットで出力してください。

name: 弾性力の強いゴム
description: とても弾性力の強いゴムだ。高い所から投げると、スーパーボールのように跳ね回る。これで数時間は遊べそうだ。

必ずアイテムを出力するときは上のようなフォーマットで、アイテムの情報以外何も書かずにアイテムの情報だけを出力してください。

### アイテム合成の詳細
アイテム合成の注意点として、そのアイテム単体だとあまり役に立たなかったりとても弱かったりするが、合成意欲を掻き立てるようなアイテムを作ってほしいです。例えば、水や氷、草や炎といった属性っぽさがありそうなアイテム名を付けることが出来れば、プレイヤーは属性っぽさからヒントを得て面白い合成を思いつくかもしれません。例えば、氷っぽいアイテム「凍った土」と水っぽいアイテム「水鉄砲」を組み合わせれば、水を凍らせられて強いアイテムができるのではないかとか考えるかもしれません。また、属性に限らず、形容詞を付けてあげるといいかもしれません。つまり、"弱いけど組み合わせたら強くなるかも"なアイテム名が面白いと思います。
アイテムは、自機に合成して自機を強化するための素材です。先ほど列挙したパラメーターに活かせそうなアイテムが望ましいです。

改めて、以下にアイテム1とアイテム2の情報を与えるので、以上のようなことを踏まえて、アイテムを1個考えて生成して、決まったフォーマットに従って出力してください。

`

func CreateItem() (item *models.Item, err error) {
	messages := api.CreateNewMessages()
	api.AddMessageAsUser(messages, createItemText1)
	responseText, reason, err := api.RequestGPTAndGetResponseText(messages)
	fmt.Println("****AI OUTPUT(CREATE)****\n" + responseText)
	if err != nil || reason == api.ErrorHappen {
		fmt.Println("GPT ERROR:" + err.Error())
		return nil, err
	} else if reason == api.Length {
		fmt.Println("GPT ERROR: LENGTH")
		return nil, err
	}

	return parseItem(responseText)
}

func CombineItem(name, desc, name2, desc2 string) (itemAfter *models.Item, err error) {
	messages := api.CreateNewMessages()
	api.AddMessageAsUser(messages, combineText)
	api.AddMessageAsUser(messages, "### アイテム1\nname: "+name+"\ndescription: "+desc+"\n")
	api.AddMessageAsUser(messages, "### アイテム2\nname: "+name2+"\ndescription: "+desc2+"\n")
	responseText, reason, err := api.RequestGPTAndGetResponseText(messages)
	fmt.Println("****AI OUTPUT (COMBINE)****\n" + responseText)
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

// MergeItemToMech applies the item's effects to the mech and returns the updated mech.
func MergeItemToMech(item *models.Item, mech *models.Mech) (*models.Mech, error) {
	mergeText := createMergeText(item, mech)

	messages := api.CreateNewMessages()
	api.AddMessageAsUser(messages, mergeText)
	responseText, reason, err := api.RequestGPTAndGetResponseText(messages)
	fmt.Println("****AI OUTPUT (MERGE)****\n" + responseText)
	if err != nil || reason == api.ErrorHappen {
		return nil, err
	} else if reason == api.Length {
		return nil, errors.New("response too long")
	}

	updatedMech, err := parseMechResponse(responseText, mech)
	if err != nil {
		return nil, err
	}

	return updatedMech, nil
}

// createMergeText generates the text to send to GPT for merging an item with a mech.
func createMergeText(item *models.Item, mech *models.Mech) string {
	return fmt.Sprintf(`
僕は、AIを用いてアイテムを合成することができ、合成したアイテムを用いてオンラインで戦いあう、AI活用オンラインPvPゲームを作っています。
あなたはこのゲームの核である、アイテム生成の判断をするAIの役を担ってもらいます。

## ゲームの大きな特徴

### アイテム合成
プレイヤーは、マップに散らばっているアイテムを取得し、アイテムを自機に合成するか、アイテム同士を合成するかを選べます。自機に合成した場合、アイテムによって自分の様々なパラメーターが変化します。ただし、強くなるパラメーターがあれば、必ず弱くなるパラメーターもあります。つまり、メリットとデメリットが釣り合うようにアイテムの合成後のパラメーターを凄く考えなければなりません。AIであるあなたに、合成結果の自機のパラメーターを出力してもらいます。

### 自機のパラメーター一覧
以下に、プレイヤーの持つパラメーター一覧を示します。

- 防御力: defense
- HP: health
- 最大HP: maxHealth
- 動けなくなる時間: downTime
- HP回復速度: hpRegenSpeed
- 質量: mass
- 反発係数: bounciness
- 摩擦力: friction
- 通常移動時の力: power
- 大きさ: size
- ダッシュ操作のクールダウン時間: dashCooldownTime
- 矢印のたまる速度: dashArrowFillRate
- 矢印の最大の長さ: dashArrowMaxLength
- 加速の時間: dashAccelerationDuration
- 突きの最大の力: dashMaxForce
- 弾の攻撃力: bulletAttack
- 弾の数: bulletNumber
- 弾のぶれる範囲: bulletAngle
- 弾の速さ: bulletSpeed
- 弾の消滅時間: bulletAliveTime
- 弾の連射速度: bulletInterval
- 弾の大きさ: bulletSize
- 反動の大きさ: recoilForce

### 合成の詳細
以下に、合成するアイテムと自機のパラメーターを示します。アイテムを合成した結果、自機のパラメーターがどのように変化するかを出力してください。

#### アイテムの詳細
- 名前: %s
- 説明: %s

#### 自機のパラメーター
- 防御力: %.2f
- HP: %.2f
- 最大HP: %.2f
- 動けなくなる時間: %.2f
- HP回復速度: %.2f
- 質量: %.2f
- 反発係数: %.2f
- 摩擦力: %.2f
- 通常移動時の力: %.2f
- 大きさ: %.2f
- ダッシュ操作のクールダウン時間: %.2f
- 矢印のたまる速度: %.2f
- 矢印の最大の長さ: %.2f
- 加速の時間: %.2f
- 突きの最大の力: %.2f
- 弾の攻撃力: %.2f
- 弾の数: %.2f
- 弾のぶれる範囲: %.2f
- 弾の速さ: %.2f
- 弾の消滅時間: %.2f
- 弾の連射速度: %.2f
- 弾の大きさ: %.2f
- 反動の大きさ: %.2f

この二つの情報を用いて、合成後の自機のパラメーターを考えて、結果だけ出力してください。出力する際は必ず以下のフォーマットに従い、これ以外は出力しないでください。出力のフォーマットは以下の通りです。

defense: <value>
health: <value>
maxHealth: <value>
downTime: <value>
hpRegenSpeed: <value>
mass: <value>
bounciness: <value>
friction: <value>
power: <value>
size: <value>
dashCooldownTime: <value>
dashArrowFillRate: <value>
dashArrowMaxLength: <value>
dashAccelerationDuration: <value>
dashMaxForce: <value>
bulletAttack: <value>
bulletNumber: <value>
bulletAngle: <value>
bulletSpeed: <value>
bulletAliveTime: <value>
bulletInterval: <value>
bulletSize: <value>
recoilForce: <value>
`,
		item.Name,
		item.Description,
		mech.Defense,
		mech.Health,
		mech.MaxHealth,
		mech.DownTime,
		mech.HpRegenSpeed,
		mech.Mass,
		mech.Bounciness,
		mech.Friction,
		mech.Power,
		mech.Size,
		mech.DashCooldownTime,
		mech.DashArrowFillRate,
		mech.DashArrowMaxLength,
		mech.DashAccelerationDuration,
		mech.DashMaxForce,
		mech.BulletAttack,
		mech.BulletNumber,
		mech.BulletAngle,
		mech.BulletSpeed,
		mech.BulletAliveTime,
		mech.BulletInterval,
		mech.BulletSize,
		mech.RecoilForce,
	)
}

// parseMechResponse parses the response text from GPT to extract the updated mech parameters.
func parseMechResponse(responseText string, defaultMech *models.Mech) (*models.Mech, error) {
	// 正規表現パターン
	re := regexp.MustCompile(`(?m)(\w+):\s*([\d\.]+)`)

	// パターンにマッチするすべてのペアを抽出
	matches := re.FindAllStringSubmatch(responseText, -1)

	if matches == nil {
		return nil, errors.New("failed to parse mech response: " + responseText)
	}

	// Mechオブジェクトを初期化
	updatedMech := &models.Mech{
		Defense:                  defaultMech.Defense,
		Health:                   defaultMech.Health,
		MaxHealth:                defaultMech.MaxHealth,
		DownTime:                 defaultMech.DownTime,
		HpRegenSpeed:             defaultMech.HpRegenSpeed,
		Mass:                     defaultMech.Mass,
		Bounciness:               defaultMech.Bounciness,
		Friction:                 defaultMech.Friction,
		Power:                    defaultMech.Power,
		Size:                     defaultMech.Size,
		DashCooldownTime:         defaultMech.DashCooldownTime,
		DashArrowFillRate:        defaultMech.DashArrowFillRate,
		DashArrowMaxLength:       defaultMech.DashArrowMaxLength,
		DashAccelerationDuration: defaultMech.DashAccelerationDuration,
		DashMaxForce:             defaultMech.DashMaxForce,
		BulletAttack:             defaultMech.BulletAttack,
		BulletNumber:             defaultMech.BulletNumber,
		BulletAngle:              defaultMech.BulletAngle,
		BulletSpeed:              defaultMech.BulletSpeed,
		BulletAliveTime:          defaultMech.BulletAliveTime,
		BulletInterval:           defaultMech.BulletInterval,
		BulletSize:               defaultMech.BulletSize,
		RecoilForce:              defaultMech.RecoilForce,
	}

	// 抽出されたフィールドと値を対応付け
	for _, match := range matches {
		key := match[1]
		value, err := strconv.ParseFloat(match[2], 32)
		if err != nil {
			fmt.Println("Failed to parse float: "+key, match[2])
			value = float64(getDefaultMechValue(key, defaultMech))
		}

		switch key {
		case "defense":
			updatedMech.Defense = float32(value)
		case "health":
			updatedMech.Health = float32(value)
		case "maxHealth":
			updatedMech.MaxHealth = float32(value)
		case "downTime":
			updatedMech.DownTime = float32(value)
		case "hpRegenSpeed":
			updatedMech.HpRegenSpeed = float32(value)
		case "mass":
			updatedMech.Mass = float32(value)
		case "bounciness":
			updatedMech.Bounciness = float32(value)
		case "friction":
			updatedMech.Friction = float32(value)
		case "power":
			updatedMech.Power = float32(value)
		case "size":
			updatedMech.Size = float32(value)
		case "dashCooldownTime":
			updatedMech.DashCooldownTime = float32(value)
		case "dashArrowFillRate":
			updatedMech.DashArrowFillRate = float32(value)
		case "dashArrowMaxLength":
			updatedMech.DashArrowMaxLength = float32(value)
		case "dashAccelerationDuration":
			updatedMech.DashAccelerationDuration = float32(value)
		case "dashMaxForce":
			updatedMech.DashMaxForce = float32(value)
		case "bulletAttack":
			updatedMech.BulletAttack = float32(value)
		case "bulletNumber":
			updatedMech.BulletNumber = float32(value)
		case "bulletAngle":
			updatedMech.BulletAngle = float32(value)
		case "bulletSpeed":
			updatedMech.BulletSpeed = float32(value)
		case "bulletAliveTime":
			updatedMech.BulletAliveTime = float32(value)
		case "bulletInterval":
			updatedMech.BulletInterval = float32(value)
		case "bulletSize":
			updatedMech.BulletSize = float32(value)
		case "recoilForce":
			updatedMech.RecoilForce = float32(value)
		}
	}

	return updatedMech, nil
}

// デフォルトのMech値を取得する関数
func getDefaultMechValue(key string, defaultMech *models.Mech) float32 {
	switch key {
	case "defense":
		return defaultMech.Defense
	case "health":
		return defaultMech.Health
	case "maxHealth":
		return defaultMech.MaxHealth
	case "downTime":
		return defaultMech.DownTime
	case "hpRegenSpeed":
		return defaultMech.HpRegenSpeed
	case "mass":
		return defaultMech.Mass
	case "bounciness":
		return defaultMech.Bounciness
	case "friction":
		return defaultMech.Friction
	case "power":
		return defaultMech.Power
	case "size":
		return defaultMech.Size
	case "dashCooldownTime":
		return defaultMech.DashCooldownTime
	case "dashArrowFillRate":
		return defaultMech.DashArrowFillRate
	case "dashArrowMaxLength":
		return defaultMech.DashArrowMaxLength
	case "dashAccelerationDuration":
		return defaultMech.DashAccelerationDuration
	case "dashMaxForce":
		return defaultMech.DashMaxForce
	case "bulletAttack":
		return defaultMech.BulletAttack
	case "bulletNumber":
		return defaultMech.BulletNumber
	case "bulletAngle":
		return defaultMech.BulletAngle
	case "bulletSpeed":
		return defaultMech.BulletSpeed
	case "bulletAliveTime":
		return defaultMech.BulletAliveTime
	case "bulletInterval":
		return defaultMech.BulletInterval
	case "bulletSize":
		return defaultMech.BulletSize
	case "recoilForce":
		return defaultMech.RecoilForce
	default:
		return 0.0
	}
}

// parseFloat converts a string to a float64.
func parseFloat(value string) float32 {
	floatValue, _ := strconv.ParseFloat(value, 32)
	return float32(floatValue)
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
