{{range $index, $elem := .todos}}
<a href="/todo/{{$elem.Id}}"><li>{{$elem.Id}} -- {{$elem.Title}}</li></a>
{{end}}

<form method="post" {{if .id}} action="/todo/{{.id}}" {{else}} action="/todo" {{end}}>
    {{if .id}} <input type="hidden" name="_method" value="PUT"> {{end}}
    <input name="title" placeholder="代办" value="{{.todo.Title}}"> <br>
    <input name="sort" type="number" placeholder="排序" value="{{.todo.Sort}}"> <br>
    <input name="status" type="number" placeholder="状态" value="{{.todo.Status}}"> <br>
    <button type="submit">提交</button>
</form>

<h3 style="color: red;">{{.message}}</h3>
