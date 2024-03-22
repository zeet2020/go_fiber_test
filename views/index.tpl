
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>media lister</title>

    <meta name="description" content="media playback">
    <meta name="author" content="media">

 

  </head>
  <body>

    <div class="container-fluid">
	<div class="row">
		<div class="col-md-12">
			<nav class="navbar navbar-expand-lg navbar-light bg-light">
				 
				<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
					<span class="navbar-toggler-icon"></span>
				</button> <a class="navbar-brand" href="#">Media Lister</a>
				<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
					
					
					
				</div>
			</nav>
		</div>
	</div>
    <div class="row">
		<div class="col-md-12">
			<table class="table">
				<thead>
					<tr>
						<th>
							#
						</th>
						<th>
							Name
						</th>
						<th>
							Path
						</th>
						<th>
							Size
						</th>
					</tr>
				</thead>
				<tbody>
                {{range $index,$val := .list }}
					<tr>
						<td>
							{{ $index }}
						</td>
						<td>
							<a href="/play/{{$val.title}}">{{ $val.title }}</a>
						</td>
						<td>
							{{ $val.url }}
						</td>
						<td>
							{{ $val.size }}
						</td>
					</tr>
					{{end}}
				</tbody>
			</table>
            
		</div>
	</div>
</div>

    
  </body>
</html>