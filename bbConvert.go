//An easy way to convert BBCode to HTML with Go.
package bbConvert

import(
    "strings"
)

//Takes in a string with BBCode and exports a string with HTML
func Convert(str string) string{
    for i:=0;i<len(str);i++{
        if str[i]=='['{
            for j:=i;j<len(str);j++{
                if str[j]==']'{
                    tmp := toHTML(str[i:j+1])
                    if tmp != str[i:j+1]{
                        str = str[:i] + tmp + str[j+1:]
                    }
                }
            }
        }
    }
    return str
}

func toHTML(str string) string{
    var beg,end string
    for i,v := range str{
        if v ==']' || v==' ' || v=='='{
            beg = str[1:i]
            break
        }
    }
    var tmp string
    for i:=len(str)-3;i>0;i--{
        if str[i:i+2]=="[/"{
            tmp = str[i:]
            end = str[i+2:len(str)-1]
            break;
        }
    }
    if beg != end{
        return str
    }
    for i,v := range str{
        if v ==']'{
            beg = str[1:i]
            break
        }
    }
    if strings.HasPrefix(tmp,"[/") && strings.HasSuffix(tmp,"]") && !isBBTag(tmp[2:len(tmp)-1]){
        return str
    }
    if len(str) - len(tmp) >1{
        str = "[" + Convert(str[1:len(str)-len(tmp)]) + tmp
    }
    str = bbToTag(str,beg)
    return str
}

func isBBTag(str string) bool{
    str = strings.ToLower(str)
    tf := str=="b"||str=="i"||str=="u"||str=="s"||str=="url"||str=="img"||str=="quote"||str=="style"||str=="color"
    return tf
}

func bbToTag(str,bb string) string{
    if bb=="img"{
        str = "<img style='float:left;width:20%;' src='" + str[5:len(str)-len(bb)] + "'/>"
    }else if strings.HasPrefix(bb,"img"){
        tagness := ""
        style := make(map[string]string)
        style["float"]="left"
        other := make(map[string]string)
        pos := make(map[string]int)
        if strings.HasPrefix(bb,"img="){
            var sz string
            for i:=5;i<len(bb);i++{
                if bb[i]==' '{
                    sz= bb[4:i]
                }else if i==len(bb)-1{
                    sz=bb[4:i+1]
                }
            }
            w,h := sz[:strings.Index(sz,"x")],sz[strings.Index(sz,"x")+1:]
            style["height"] = h
            style["width"] = w
        }
        if strings.Contains(bb,"alt=\"")||strings.Contains(bb,"alt='"){
            pos["alt"]=strings.Index(bb,"alt=")
            for i:=strings.Index(bb,"alt=")+5;i<len(bb);i++{
                if (bb[i]==bb[strings.Index(bb,"alt=")+4]&&bb[i-1]!='\\')||bb[i]==']'{
                    other["alt"]=bb[strings.Index(bb,"alt=")+5:i]
                    pos["altEnd"]=i
                    break
                }
            }
        }
        if strings.Contains(bb,"title=\"")||strings.Contains(bb,"title='"){
            pos["title"]=strings.Index(bb,"title=")
            for i:=strings.Index(bb,"title=")+7;i<len(bb);i++{
                if (bb[i]==bb[strings.Index(bb,"title=")+6]&&bb[i-1]!='\\')||bb[i]==']'{
                    other["title"]=bb[strings.Index(bb,"title=")+7:i]
                    pos["titleEnd"]=i
                    break
                }
            }
        }
        if strings.Contains(bb,"width="){
            if (pos["alt"] == 0 || strings.Index(bb,"width=") < pos["alt"]) && (pos["title"] ==0 || strings.Index(bb,"width=") < pos["title"]){
                var sz string
                for i:=strings.Index(bb,"width=")+7;i<len(bb);i++{
                    if bb[i]==' '||bb[i]=='"'||bb[i]=='\''{
                        sz= bb[strings.Index(bb,"width=")+6:i]
                        break;
                    }else if i==len(bb)-1{
                        sz=bb[strings.Index(bb,"width=")+6:i+1]
                        break;
                    }
                }
                sz = strings.Replace(sz,"\"","",-1)
                sz = strings.Replace(sz,"'","",-1)
                style["width"]=sz
            }else if (pos["altEnd"] == 0 || strings.LastIndex(bb,"width=") > pos["altEnd"]) && (pos["titleEnd"] ==0 || strings.LastIndex(bb,"width=") > pos["titleEnd"]){
                var sz string
                for i:=strings.LastIndex(bb,"width=")+7;i<len(bb);i++{
                    if bb[i]==' '||bb[i]=='"'||bb[i]=='\''{
                        sz= bb[strings.LastIndex(bb,"width=")+6:i]
                        break;
                    }else if i==len(bb)-1{
                        sz=bb[strings.LastIndex(bb,"width=")+6:i+1]
                        break;
                    }
                }
                sz = strings.Replace(sz,"\"","",-1)
                sz = strings.Replace(sz,"'","",-1)
                style["width"]=sz
            }
        }
        if strings.Contains(bb,"height="){
            if (pos["alt"]==0 || strings.Index(bb,"height=") < pos["alt"]) && (pos["title"]==0 || strings.Index(bb,"height=") < pos["title"]){
                var sz string
                for i:=strings.Index(bb,"height=")+7;i<len(bb);i++{
                    if bb[i]==' '||bb[i]=='"'||bb[i]=='\''{
                        sz= bb[strings.Index(bb,"height=")+7:i]
                        break;
                    }else if i==len(bb)-1{
                        sz=bb[strings.Index(bb,"height=")+7:i+1]
                        break;
                    }
                }
                sz = strings.Replace(sz,"\"","",-1)
                sz = strings.Replace(sz,"'","",-1)
                style["height"]=sz
            }else if (pos["altEnd"]==0 || strings.LastIndex(bb,"height=") > pos["altEnd"]) && (pos["titleEnd"]==0 || strings.LastIndex(bb,"height=") > pos["titleEnd"]){
                var sz string
                for i:=strings.LastIndex(bb,"height=")+7;i<len(bb);i++{
                    if bb[i]==' '||bb[i]=='"'||bb[i]=='\''{
                        sz= bb[strings.LastIndex(bb,"height=")+7:i]
                        break;
                    }else if i==len(bb)-1{
                        sz=bb[strings.LastIndex(bb,"height=")+7:i+1]
                        break;
                    }
                }
                sz = strings.Replace(sz,"\"","",-1)
                sz = strings.Replace(sz,"'","",-1)
                style["height"]=sz
            }
        }
        if strings.Contains(bb,"left"){
            if ((pos["alt"]==0 || strings.Index(bb,"left") < pos["alt"]) && (pos["title"]==0 || strings.Index(bb,"left") < pos["title"])) || ((pos["altEnd"]==0 || strings.LastIndex(bb,"left") > pos["altEnd"]) && (pos["titleEnd"]==0 || strings.LastIndex(bb,"left") > pos["titleEnd"])){
                style["float"]="left"
            }
        }else if strings.Contains(bb,"right"){
            if ((pos["alt"]==0 || strings.Index(bb,"right") < pos["alt"]) && (pos["title"]==0 || strings.Index(bb,"right") < pos["title"])) || ((pos["altEnd"]==0 || strings.LastIndex(bb,"right") > pos["altEnd"]) && (pos["titleEnd"]==0 || strings.LastIndex(bb,"right") > pos["titleEnd"])){
                style["float"]="right"
            }
        }
        tagness = " style='"
        for i,v := range style{
            tagness += i + ":" + v + ";"
        }
        tagness += "'"
        if other["alt"]!=""{
            tagness += " alt='"+other["alt"]+"'"
        }
        if other["title"]!=""{
            tagness += " title='"+other["title"]+"'"
        }
        str = "<img"+tagness+" src='"+str[len(bb)+2:len(str)-6]+"'/>"
    }else if bb=="b" || bb=="i" || bb=="u" || bb=="s"{
        str = "<"+bb+">"+str[3:len(str)-4]+"</"+bb+">"
    }else if bb=="url"{
        str = "<a href='" + str[5:len(str)-6] + "'>" + str[5:len(str)-6] + "</a>"
    }else if strings.HasPrefix(bb,"url="){
        str = "<a href='" + bb[5:]+"'>" + str[len(bb)+2:len(str)-6] + "</a>"
    }else if strings.HasPrefix(bb,"color="){
        str = "<span style='color:" + bb[7:] + ";'>" + str[len(bb)+2:len(str)-8] + "</span>"
    }else if strings.HasPrefix(bb,"quote=\"")|| strings.HasPrefix(bb,"quote='"){
        str = "<div class='quote'>"+bb[7:len(bb)-1]+"<blockquote>"+str[len(bb)+2:len(str)-8]+"</blockquote></div>"
    }else if strings.HasPrefix(bb,"quote="){
        str = "<div class='quote'>"+bb[6:]+"<blockquote>"+str[len(bb)+2:len(str)-8]+"</blockquote></div>"
    }
    return str
}