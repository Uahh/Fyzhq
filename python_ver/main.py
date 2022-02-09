import sys
import json
import pypinyin

sentence = pypinyin.pinyin(
    '演的吧。演的怎么了？人生如戏，不能演吗？真搞不懂为什么有些人总是天天在下面怀疑别人演不演的，爱看不看不看拉倒！就是演的也比你们强，整天看到你们在这里怀疑这个怀疑那个，真的气的我眼泪在眼眶打转，手攥紧了衣角整个人都在颤抖，太让我失望了更多的是心寒，我想我需要退网一段时间冷静冷静，希望大家支持。', style=pypinyin.NORMAL)

with open("./data/language.json", encoding='utf_8') as json_file:
    language = json.load(json_file)
    json_file.close()

result = ""

if sys.argv[1] == "-jp":
    for word in sentence:
        result += language["Japanese"][word[0]]
elif sys.argv[1] == "-gm":
    for word in sentence:
        result += language["German"][word[0]]
elif sys.argv[1] == "-fr":
    for word in sentence:
        result += language["French"][word[0]]
elif sys.argv[1] == "-ru":
    for word in sentence:
        result += language["Russian"][word[0]]
elif sys.argv[1] == "-kr":
    for word in sentence:
        result += language["Korean"][word[0]]
elif sys.argv[1] == "-th":
    for word in sentence:
        result += language["Thai"][word[0]]
else:
    result = "Error"

print(result)
