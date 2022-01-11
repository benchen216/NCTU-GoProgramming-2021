import re
ans1 = '''  -max int
    	Max Printing (default 10)
  -w string
    	Web page (default "ptt")

'''
ans2 = '''flag needs an argument: -w
Usage of:
  -max int
    	Max Printing (default 10)
  -w string
    	Web page (default "ptt")
exit status 2

'''

ans3 = '''flag needs an argument: -max
Usage of:
  -max int
    	Max Printing (default 10)
  -w string
    	Web page (default "ptt")
exit status 2

'''

ans4 = '''1. 名字: LineageM, 留言: 讚 禮拜一漲就是強 跌就是套, 時間: 01/08 18:35
2. 名字: b23058179, 留言: 39被甩下去QQ, 時間: 01/08 18:35
3. 名字: iDoDo, 留言: 丸子, 時間: 01/08 18:35
4. 名字: ntupeap, 留言: 樓下支援v7馬頭工人, 時間: 01/08 18:35
5. 名字: ping860622, 留言: 新聞可以不要那麼快出嗎, 時間: 01/08 18:35

'''

ans5 = '''1. 姓名: 張燕光, 網站: http://cial.csie.ncku.edu.tw/
2. 姓名: 王士豪, 網站: NULL
3. 姓名: 陳朝鈞, 網站: http://nckunetdb.appspot.com/
4. 姓名: 李同益, 網站: http://graphics.csie.ncku.edu.tw/
5. 姓名: 吳宗憲, 網站: http://chinese.csie.ncku.edu.tw/
6. 姓名: 鄭芳田, 網站: http://140.116.234.1:3030/
7. 姓名: 謝孫源, 網站: http://algorithm.csie.ncku.edu.tw/
8. 姓名: 孫永年, 網站: http://sites.google.com/view/ncku-csie-vslab
9. 姓名: 黃崇明, 網站: http://www.mmnetlab.csie.ncku.edu.tw/
10. 姓名: 黃宗立, 網站: http://myweb.ncku.edu.tw/~islab62524/


'''

ans6 = '''1. 姓名: 張燕光, 網站: http://cial.csie.ncku.edu.tw/
2. 姓名: 王士豪, 網站: NULL
3. 姓名: 陳朝鈞, 網站: http://nckunetdb.appspot.com/

'''

ans7 = '''1. 姓名: 張燕光, 網站: http://cial.csie.ncku.edu.tw/
2. 姓名: 王士豪, 網站: NULL
3. 姓名: 陳朝鈞, 網站: http://nckunetdb.appspot.com/
4. 姓名: 李同益, 網站: http://graphics.csie.ncku.edu.tw/
5. 姓名: 吳宗憲, 網站: http://chinese.csie.ncku.edu.tw/
6. 姓名: 鄭芳田, 網站: http://140.116.234.1:3030/
7. 姓名: 謝孫源, 網站: http://algorithm.csie.ncku.edu.tw/
8. 姓名: 孫永年, 網站: http://sites.google.com/view/ncku-csie-vslab
9. 姓名: 黃崇明, 網站: http://www.mmnetlab.csie.ncku.edu.tw/
10. 姓名: 黃宗立, 網站: http://myweb.ncku.edu.tw/~islab62524/
11. 姓名: 郭耀煌, 網站: http://ismp.csie.ncku.edu.tw/
12. 姓名: 陳裕民, 網站: https://sblab.imis.ncku.edu.tw/
13. 姓名: 蔣榮先, 網站: http://iir.csie.ncku.edu.tw/
14. 姓名: 陳培殷, 網站: http://dic.csie.ncku.edu.tw/
15. 姓名: 李強, 網站: http://dblab.csie.ncku.edu.tw/home/index.html
16. 姓名: 陳響亮, 網站: http://140.116.86.195/
17. 姓名: 鄭憲宗, 網站: http://plato.csie.ncku.edu.tw/
18. 姓名: 楊大和, 網站: https://mmlab.imis.ncku.edu.tw/
19. 姓名: 蘇文鈺, 網站: http://screamlab-ncku-2008.blogspot.com/
20. 姓名: 郭淑美, 網站: http://idip.csie.ncku.edu.tw
21. 姓名: 連震杰, 網站: http://robotics.csie.ncku.edu.tw/
22. 姓名: 蘇銓清, 網站: http://140.116.247.66/
23. 姓名: 蕭宏章, 網站: http://140.116.246.157
24. 姓名: 高宏宇, 網站: https://ikmlab.csie.ncku.edu.tw
25. 姓名: 盧文祥, 網站: http://wmmks.csie.ncku.edu.tw/wmmks/
26. 姓名: 梁勝富, 網站: http://ncbci.csie.ncku.edu.tw/
27. 姓名: 張大緯, 網站: http://os.csie.ncku.edu.tw/
28. 姓名: 藍崑展, 網站: https://lens.csie.ncku.edu.tw
29. 姓名: 林英超, 網站: http://caid.csie.ncku.edu.tw/
30. 姓名: 賀保羅, 網站: NULL
31. 姓名: 朱威達, 網站: http://mmcv.csie.ncku.edu.tw
32. 姓名: 蔡孟勳, 網站: http://imslab.org/
33. 姓名: 許靜芳, 網站: http://hsnl.csie.ncku.edu.tw/
34. 姓名: 楊中平, 網站: http://neat.csie.ncku.edu.tw
35. 姓名: 吳明龍, 網站: http://bmilab.csie.ncku.edu.tw
36. 姓名: 莊坤達, 網站: https://ncku-ccs.github.io/netdb-web/
37. 姓名: 蔡佩璇, 網站: http://cps.imis.ncku.edu.tw/
38. 姓名: 涂嘉恒, 網站: http://chiaheng.wordpress.com/advanced-systems-research-lab/
39. 姓名: 陳奇業, 網站: http://sivslab.csie.ncku.edu.tw
40. 姓名: 王宏鍇, 網站: http://splab.imis.ncku.edu.tw/
41. 姓名: 曾繁勛, 網站: NULL
42. 姓名: 何建忠, 網站: NULL
43. 姓名: 李信杰, 網站: https://reurl.cc/GbYoeW
44. 姓名: 張瑞紘, 網站: https://cssa.cc.ncku.edu.tw/ladder/
45. 姓名: 黃敬群, 網站: NULL

'''

ans8 = '''1. 名字: LineageM, 留言: 讚 禮拜一漲就是強 跌就是套, 時間: 01/08 18:35
2. 名字: b23058179, 留言: 39被甩下去QQ, 時間: 01/08 18:35
3. 名字: iDoDo, 留言: 丸子, 時間: 01/08 18:35
4. 名字: ntupeap, 留言: 樓下支援v7馬頭工人, 時間: 01/08 18:35
5. 名字: ping860622, 留言: 新聞可以不要那麼快出嗎, 時間: 01/08 18:35
6. 名字: Coffeewater, 留言: 38.3砍掉的舉個手, 時間: 01/08 18:36
7. 名字: kangta2030, 留言: 幹這新聞標題, 時間: 01/08 18:36
8. 名字: Manhood27, 留言: 利多出盡，股價早已反應, 時間: 01/08 18:36
9. 名字: mtyk10100, 留言: 38被甩下去 40補回來了, 時間: 01/08 18:36
10. 名字: beforehand, 留言: 雖然心裡煎熬 還好信仰夠 沒被甩出去, 時間: 01/08 18:36
11. 名字: lizard30923, 留言: 樓下38.3市價賣韭菜, 時間: 01/08 18:36
12. 名字: lmc66, 留言: 又要連漲停兩天然後洗盤了, 時間: 01/08 18:36
13. 名字: chingkai, 留言: 爽啦！F87週二記得要補票，因為週一漲停買不到, 時間: 01/08 18:37
14. 名字: zeem, 留言: 是，我是菜雞。星期一還買的到船票嗎？, 時間: 01/08 18:37
15. 名字: t73697, 留言: 記者買幾張??新聞發這麼快, 時間: 01/08 18:37
16. 名字: s155260, 留言: 馬頭工人, 時間: 01/08 18:37
17. 名字: fedona, 留言: 砍在阿呆谷的韭菜表示：, 時間: 01/08 18:37
18. 名字: evil006, 留言: 然後一月是小月變月減，二月過年＋只有28天可能也是, 時間: 01/08 18:38
19. 名字: evil006, 留言: 月減, 時間: 01/08 18:38
20. 名字: chingkai, 留言: 信仰不夠一率不要買，船有夠晃, 時間: 01/08 18:38
21. 名字: kitsune318, 留言: 38.3砍60張的阿呆+1, 時間: 01/08 18:38
22. 名字: e73103999, 留言: 哈哈哈笑死 前幾天有人說大戶偷看答案在出貨, 時間: 01/08 18:38
23. 名字: lizard30923, 留言: 12月運價漲幅很大，一月營收根本不用擔心, 時間: 01/08 18:38
24. 名字: lizard30923, 留言: 一月目前也都還在漲，二月營收應該也是穩的, 時間: 01/08 18:39
25. 名字: evil006, 留言: 昨天洗掉一半船票之後腦袋有比較清醒了，現在財報是, 時間: 01/08 18:39
26. 名字: strayfrog, 留言: 禮拜一可能會出貨，但有跌就撿, 時間: 01/08 18:39
27. 名字: evil006, 留言: 目前能聽到的最後利多, 時間: 01/08 18:39
28. 名字: laipipi, 留言: 這不是本來就反應？, 時間: 01/08 18:39
29. 名字: nidhogg, 留言: 自由記者又發了, 時間: 01/08 18:39
30. 名字: a52655, 留言: QQ被騙月增29%!,(MISSING), 時間: 01/08 18:39

'''
ans_list = ["-1",ans1,ans2,ans3,ans4,ans5,ans6,ans7,ans8]
cmd_list = ['-1',
'go run lab11.go',
'go run lab11.go -w',
'go run lab11.go -max',
'go run lab11.go -max 5',
'go run lab11.go -w ncku',
'go run lab11.go -w ncku -max 3',
'go run lab11.go -max 100 -w ncku',
'go run lab11.go -max 30 -w ptt']
count = 0
def remove_rn(arr):
    arr = arr.replace('\r', '\n')
    arr = arr.replace('\n', '')
    return arr

for i in range(1,9):
    with open('result' + str(i) + '.txt','r',encoding='utf-8') as f:
        result = f.read()
        usage_string = re.findall(r'Usage of.*:+', result)
        try:
            result = result.replace(usage_string[0], 'Usage of:')
        except:
            pass
        your_result = remove_rn(result)
        ans = remove_rn(ans_list[i])
        print('-------------------------------------------------')
        print('Action{id}: {cmd}'.format(id=i, cmd=cmd_list[i]))
        if ans == your_result:
            count += 1
            print('GET POINT 1')
        else:
            print('Your Result: \n{result}\n*****************************************************\nAns: \
                    \n{ans}Wrong Answer ; NO POINT'.format(result=result, ans=ans_list[i]))
                    
print('Pass: {num}/8'.format(num=count))
