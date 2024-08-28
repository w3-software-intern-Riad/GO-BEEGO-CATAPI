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
        <div class="cat-image">
            <!-- Placeholder image if no cat image is available -->
            <img src="" alt="Cat Image" id="cat-image">
        </div>
        <div class="actions">
            <button class="love-btn" style="font-size:24px"> <i class="fa fa-heart"></i></button>
            
            <button class="like-btn" style="font-size:24px">👍</button>
            <button class="dislike-btn" style="font-size:24px">👎</button>
        </div>
    </div>

    <script src="/static/js/main.js"></script>
</body>
</html>
