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

ans5 = '''1. 姓名: 洪士灝 (SH Hung), 個人網站: http://www.csie.ntu.edu.tw/~hungsh/
2. 姓名: 張傑帆(Jie-Fan Chang), 個人網站: http://homepage.ntu.edu.tw/~jfanc/
3. 姓名: 陳祝嵩 (Chu-Song Chen), 個人網站: http://imp.iis.sinica.edu.tw/
4. 姓名: 陳文進 (WC Chen), 個人網站: http://www.cmlab.csie.ntu.edu.tw/~wcchen/
5. 姓名: 鄭卜壬 (PJ Cheng), 個人網站: http://www.csie.ntu.edu.tw/~pjcheng/
6. 姓名: 項潔 (Jieh Hsiang), 個人網站: http://www.digital.ntu.edu.tw/hsiang/
7. 姓名: 薛智文 (CW Hsueh), 個人網站: http://www.csie.ntu.edu.tw/~cwhsueh/
8. 姓名: 賴飛羆 (FP Lai), 個人網站: https://sites.google.com/site/medinfolabatntu/
9. 姓名: 黎士瑋 (Shih-Wei Li), 個人網站: https://www.csie.ntu.edu.tw/~shihwei/
10. 姓名: 林忠緯 (CW Lin), 個人網站: https://www.csie.ntu.edu.tw/~cwlin/

'''

ans6 = '''1. 姓名: 洪士灝 (SH Hung), 個人網站: http://www.csie.ntu.edu.tw/~hungsh/
2. 姓名: 張傑帆(Jie-Fan Chang), 個人網站: http://homepage.ntu.edu.tw/~jfanc/
3. 姓名: 陳祝嵩 (Chu-Song Chen), 個人網站: http://imp.iis.sinica.edu.tw/

'''

ans7 = '''1. 姓名: 洪士灝 (SH Hung), 個人網站: http://www.csie.ntu.edu.tw/~hungsh/
2. 姓名: 張傑帆(Jie-Fan Chang), 個人網站: http://homepage.ntu.edu.tw/~jfanc/
3. 姓名: 陳祝嵩 (Chu-Song Chen), 個人網站: http://imp.iis.sinica.edu.tw/
4. 姓名: 陳文進 (WC Chen), 個人網站: http://www.cmlab.csie.ntu.edu.tw/~wcchen/
5. 姓名: 鄭卜壬 (PJ Cheng), 個人網站: http://www.csie.ntu.edu.tw/~pjcheng/
6. 姓名: 項潔 (Jieh Hsiang), 個人網站: http://www.digital.ntu.edu.tw/hsiang/
7. 姓名: 薛智文 (CW Hsueh), 個人網站: http://www.csie.ntu.edu.tw/~cwhsueh/
8. 姓名: 賴飛羆 (FP Lai), 個人網站: https://sites.google.com/site/medinfolabatntu/
9. 姓名: 黎士瑋 (Shih-Wei Li), 個人網站: https://www.csie.ntu.edu.tw/~shihwei/
10. 姓名: 林忠緯 (CW Lin), 個人網站: https://www.csie.ntu.edu.tw/~cwlin/
11. 姓名: 呂學一 (HI Lu), 個人網站: http://www.csie.ntu.edu.tw/~hil/
12. 姓名: 陳偉松 (Tony Tan), 個人網站: http://www.csie.ntu.edu.tw/~tonytan/
13. 姓名: 曾宇鳳 (YF Tseng), 個人網站: https://www.cmdm.tw/
14. 姓名: 陳炳宇 (BY Chen), 個人網站: http://graphics.csie.ntu.edu.tw/~robin/
15. 姓名: 陳尚澤 (Shang-Tse Chen), 個人網站: https://www.csie.ntu.edu.tw/~stchen/
16. 姓名: 鄭龍磻（Lung-Pan Cheng）, 個人網站: http://lungpancheng.tw
17. 姓名: 傅楸善 (CS Fuh), 個人網站: http://www.csie.ntu.edu.tw/~fuh/
18. 姓名: 徐宏民 (Winston Hsu), 個人網站: http://winstonhsu.info/
19. 姓名: 郭大維 (TW Kuo), 個人網站: http://www.csie.ntu.edu.tw/%7Ektw/
20. 姓名: 李明穗 (MS Lee), 個人網站: http://www.csie.ntu.edu.tw/~mslee/
21. 姓名: 林智仁 (CJ Lin), 個人網站: http://www.csie.ntu.edu.tw/~cjlin/
22. 姓名: 劉邦鋒 (PF Liu), 個人網站: http://www.csie.ntu.edu.tw/~pangfeng/
23. 姓名: 逄愛君 (AC Pang), 個人網站: http://www.csie.ntu.edu.tw/%7Eacpang/
24. 姓名: 楊佳玲 (CL Yang), 個人網站: NULL
25. 姓名: 李德財 (DT Lee), 個人網站: http://www.iis.sinica.edu.tw/~dtlee/
26. 姓名: 傅立成 (LC Fu), 個人網站: http://robotlab.csie.ntu.edu.tw/
27. 姓名: 張韻詩 (W.-S. Chang), 個人網站: http://www.iis.sinica.edu.tw/pages/janeliu/index_en.html
28. 姓名: 馬匡六 (Kwan-Liu Ma), 個人網站: http://www.cs.ucdavis.edu/~ma/
29. 姓名: 卓政宏 (CH Cho), 個人網站: NULL
30. 姓名: 徐慰中 (WC Hsu), 個人網站: https://covart.csie.ntu.edu.tw/advisor
31. 姓名: 王柏堯 (BY Wang), 個人網站: NULL
32. 姓名: 徐讚昇 (TS Hsu), 個人網站: http://homepage.iis.sinica.edu.tw/~tshsu/
33. 姓名: 歐陽明 (M Ouhyoung), 個人網站: http://www.csie.ntu.edu.tw/~ming

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
30. 名字: a52655, 留言: QQ被騙月增29%, 時間: 01/08 18:39

'''
ans_list = ["-1",ans1,ans2,ans3,ans4,ans5,ans6,ans7,ans8]
cmd_list = ['-1',
'go run lab11.go',
'go run lab11.go -w',
'go run lab11.go -max',
'go run lab11.go -max 5',
'go run lab11.go -w ntu',
'go run lab11.go -w ntu -max 3',
'go run lab11.go -max 100 -w ntu',
'go run lab11.go -max 30 -w ptt']
count = 0
def remove_rn(arr):
    arr = arr.replace('\r', '\n')
    arr = arr.replace('\n', '')
    return arr

for i in range(1,9):
    with open('result' + str(i) + '.txt') as f:
        result = f.read()
        usage_string = re.findall(r'Usage.*?:', result)
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
