<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Voting App</title>
    <link rel="stylesheet" href="/static/css/style.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    
</head>
<body>
    <div class="container">
         <div class="header">
            <button class="nav-btn" onclick="window.location.href='/';">Voting</button>
            <button class="nav-btn" onclick="window.location.href='/get-breeds';">Breeds</button>
            <button class="nav-btn" onclick="window.location.href='/get-favorite';">Favs</button>
        </div>

        <div id="favorites-container" class="favorites-container">
            
        </div>
        
    </div>

    <script src="/static/js/getFav.js"></script>
</body>
</html>
