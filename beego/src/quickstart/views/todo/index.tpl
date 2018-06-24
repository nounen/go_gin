{{range $index, $elem := .todos}}
    <li>{{$elem.Id}} -- {{$elem.Title}}</li>
{{end}}

<form method="post" action="todo">
    <input name="title" placeholder="代办" value="默认代办"> <br>
    <input name="sort" placeholder="排序" value="0"> <br>
    <input name="status" placeholder="状态" value="1"> <br>
    <button type="submit">提交</button>
</form>

<h3 style="color: red;">{{.message}}</h3>
