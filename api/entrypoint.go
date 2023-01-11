package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/chess", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
<!DOCTYPE html>
<html>

    <head>
        <style>
            * {
                margin: 0;
                padding: 0;
                transition: 0.2s ease;
            }

            body {
                display: flex;
                justify-content: center;
                align-items: center;
                width: 100vw;
                height: 100vh;
            }

            #board {
                outline: 4px solid black;
            }

            .box {
                display: inline-flex;
                width: 8vh;
                height: 8vh;
                border: 1px solid black;
                justify-content: center;
                align-items: center;
                cursor: pointer;
            }

            #blackkill p,
            #whitekill p {
                font-size: 3em;
                letter-spacing: -40px;
                text-shadow: -1px -1px 0px black;
            }

            #blackkill,
            #whitekill {
                width: 16vh;
                height: 65vh;
                color: goldenrod;
            }

            #blackkill {
                color: darkgoldenrod;
            }

            .row {
                margin-bottom: -4px;
            }

            .row:nth-child(odd) .box:nth-child(odd),
            .row:nth-child(even) .box:nth-child(even) {
                background-color: darkslategrey;
                filter: brightness(1.3);
            }

            .row:nth-child(odd) .box:nth-child(even),
            .row:nth-child(even) .box:nth-child(odd) {
                background-color: #C2C2C2;
                filter: brightness(1.3);
            }

            p.p {
                font-size: 9.5vh;
                color: goldenrod;
                text-shadow: 1px 1px 1px black;
            }

            .black p {
                color: darkgoldenrod;
            }

            #winner {
                width: 100vw;
                height: 100vh;
                position: absolute;
                background-color: rgba(0, 0, 0, 0.8);
                display: none;
                justify-content: center;
                align-items: center;
            }

            #winner p {
                font-size: 3em;
                color: whitesmoke;
            }
        </style>
    </head>

    <body>
        <script>
            window.onload = function () {
                board();
                spices();
            }
            var html = "",
                a = "its alive!!!",
                bs = Math.min(window.innerHeight - 20, window.innerWidth - 20),
                block = bs / 8,
                wp = 9817,
                wk = 9812,
                wq = 9813,
                wr = 9814,
                wb = 9815,
                wh = 9816,
                bk = 9818,
                bq = 9819,
                br = 9820,
                bb = 9821,
                bh = 9822,
                bp = 9823,
                id = "",
                place = 0,
                pcolor = "",
                pname = "",
                pturn = "",
                lpcolor = "",
                lpname = "",
                lpturn = "",
                j = 0,
                p = "",
                arry = [],
                cls = "",
                lastval = 0,
                turn = ["turn", "nturn"],
                turnval = 0,
                lastp = "";
            function spices() {
                //  black pawn
                for (var i = 9; i <= 16; i++) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + bp + ";</p>";
                    document.getElementById("b" + i).className = "box black pawn nturn";
                }
                //  black rook
                for (var i = 1; i <= 8; i += 7) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + br + ";</p>";
                    document.getElementById("b" + i).className = "box black rook nturn";
                }
                //  black horse
                for (var i = 2; i <= 7; i += 5) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + bh + ";</p>";
                    document.getElementById("b" + i).className = "box black horse nturn";
                }
                //  black biship
                for (var i = 3; i <= 6; i += 3) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + bb + ";</p>";
                    document.getElementById("b" + i).className = "box black biship nturn";
                }
                //  black queen
                for (var i = 5; i <= 5; i++) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + bq + ";</p>";
                    document.getElementById("b" + i).className = "box black queen nturn";
                }
                //  black king
                for (var i = 4; i <= 4; i++) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + bk + ";</p>";
                    document.getElementById("b" + i).className = "box black king nturn";
                }
                //  white rook
                for (var i = 8 * 7; i > 8 * 6; i--) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + wp + ";</p>";
                    document.getElementById("b" + i).className = "box white pawn turn";
                }
                //  white rook
                for (var i = 8 * 8; i > 8 * 7; i -= 7) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + wr + ";</p>";
                    document.getElementById("b" + i).className = "box white rook turn";
                }
                //  white horse
                for (var i = 63; i > 56; i -= 5) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + wh + ";</p>";
                    document.getElementById("b" + i).className = "box white horse turn";
                }
                //  white biship
                for (var i = 62; i > 56; i -= 3) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + wb + ";</p>";
                    document.getElementById("b" + i).className = "box white biship turn";
                }
                //  white queen
                for (var i = 61; i >= 61; i--) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + wq + ";</p>";
                    document.getElementById("b" + i).className = "box white queen turn";
                }
                //  white king
                for (var i = 60; i >= 60; i--) {
                    document.getElementById("b" + i).innerHTML = "<p class='p'>&#" + wk + ";</p>";
                    document.getElementById("b" + i).className = "box white king turn";
                }
            }
            function board() {
                for (var i = 0; i < 8; i++) {
                    html += "<div class='row'>";
                    for (var j = 1; j <= 8; j++) {
                        html += "<div class='box pcolor none nturn' id='b" + (i * 8 + j) + "' onclick=pice(" + (i * 8 + j) + ")></div>";
                    }
                    html += "</div>";
                }
                document.getElementById("board").innerHTML = html;
            }
            function pice(val) {
                id = document.getElementById("b" + val);
                var j = 0; //for cell styling
                //clearing other paths
                for (var i = 1; i <= 8 * 8; i++) {
                    document.getElementById("b" + i).style.filter = "brightness(1.3) sepia(0)"
                }
                //cheching empty cells
                if (id.innerHTML != "") {
                    pcolor = id.classList[1];
                    pname = id.classList[2];
                    pturn = id.classList[3];
                    p = id.innerHTML;
                    //checking black pice
                    if (pcolor == "black" && turn[turnval] == pturn) {
                        arry = [];
                        //  black pawn move
                        if (pname == "pawn") {
                            lastp = p;
                            //for starting position
                            if (val > 8 && val <= 16) {
                                //looping to get all three cells at once
                                for (var i = 0; i < 2; i++) {
                                    j += 8;
                                    document.getElementById("b" + val).style.filter = "brightness(1)";
                                    //checking next cell is empty
                                    if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                        document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                        arry.push(val + j);
                                    }
                                } //end of for loop
                            } else {
                                j += 8
                                document.getElementById("b" + val).style.filter = "brightness(1)";
                                //checking next cell
                                if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                }
                            }
                            //checking none last pawn
                            if (val % 8 != 0 && document.getElementById("b" + (val + 9)).classList[1] == "white") {
                                document.getElementById("b" + (val + 9)).style.filter = "sepia(1)";
                                arry.push(val + 9);
                            }
                            //checking none fist pawn
                            if ((val - 1) % 8 != 0 && document.getElementById("b" + (val + 7)).classList[1] == "white") {
                                document.getElementById("b" + (val + 7)).style.filter = "sepia(1)";
                                arry.push(val + 7);
                            }
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                        } //end of pawn if
                        //black rook move
                        if (pname == "rook" || pname == "queen") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            //bottom move
                            j = 8;
                            var bottommove = 0;
                            var test = val;
                            for (var i = 0; i < 8; i++) {
                                test += 8
                                if (test < 65) {
                                    bottommove += 1;
                                }
                            }
                            for (var i = 0; i < bottommove; i++) {
                                document.getElementById("b" + val).style.filter = "brightness(1)";
                                if (document.getElementById("b" + (val + j)).classList[1] == "white") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                //checking for empty cell
                                if (document.getElementById("b" + (val + j)).classList[2] == "none" && (val + j) < 65) {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j += 8;
                                }
                            } //end of bottom move
                            //top move
                            j = -8;
                            var topmove = 0;
                            test = val
                            for (var i = 0; i < 8; i++) {
                                test -= 8
                                if (test > 0) {
                                    topmove += 1;
                                }
                            }
                            for (var i = 0; i < topmove; i++) {
                                if (document.getElementById("b" + (val + j)).classList[1] == "white") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                //checking for empty cells
                                if ((val + j) > 0 && document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j -= 8;
                                }
                            } //end of top move
                            //right move
                            j = 1;
                            var moveright = 0;
                            //checking not at right edge
                            if (val % 8 == 0) {
                                moveright = 0
                            } else {
                                moveright = 8 - (val % 8);
                            }
                            for (var i = 0; i < moveright; i++) {
                                if (document.getElementById("b" + (val + j)).classList[1] == "white") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j += 1;
                                }
                            } //end of right move
                            //left move{
                            j = -1;
                            //checking not at left edge
                            var moveleft = 0;
                            if (val % 8 == 0) {
                                moveleft = 7;
                            } else {
                                moveleft = (val - 1) % 8;
                            }
                            for (var i = 0; i < moveleft; i++) {
                                if (document.getElementById("b" + (val + j)).classList[1] == "white") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j -= 1;
                                }
                            }
                        } //end of rook
                        //black horse
                        if (pname == "horse") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            var leftbottom = val + 8 + 7;
                            var rightbottom = val + 8 + 9;
                            var bottomright = val + 10;
                            var bottomleft = val + 6;
                            var righttop = val - 8 - 7;
                            var lefttop = val - 8 - 9;
                            var topright = val - 6;
                            var topleft = val - 10;
                            document.getElementById("b" + val).style.filter = "brightness(1)";
                            if (leftbottom < 65 && leftbottom % 8 != 0) {
                                if (document.getElementById("b" + leftbottom).classList[2] == "none") {
                                    document.getElementById("b" + leftbottom).style.filter = "brightness(1)";
                                    arry.push(leftbottom);
                                }
                                if (document.getElementById("b" + leftbottom).classList[1] == "white") {
                                    document.getElementById("b" + leftbottom).style.filter = "sepia(1)";
                                    arry.push(leftbottom);
                                }
                            }
                            if (rightbottom < 65 && val % 8 != 0) {
                                if (document.getElementById("b" + rightbottom).classList[2] == "none") {
                                    document.getElementById("b" + rightbottom).style.filter = "brightness(1)";
                                    arry.push(rightbottom);
                                }
                                if (document.getElementById("b" + rightbottom).classList[1] == "white") {
                                    document.getElementById("b" + rightbottom).style.filter = "sepia(1)";
                                    arry.push(rightbottom);
                                }
                            }
                            if (bottomleft < 65 && (val - 1) % 8 != 0 && (val - 2) % 8 != 0) {
                                if (document.getElementById("b" + bottomleft).classList[2] == "none") {
                                    document.getElementById("b" + bottomleft).style.filter = "brightness(1)";
                                    arry.push(bottomleft);
                                }
                                if (document.getElementById("b" + bottomleft).classList[1] == "white") {
                                    document.getElementById("b" + bottomleft).style.filter = "sepia(1)";
                                    arry.push(bottomleft);
                                }
                            }
                            if (bottomright < 65 && val % 8 != 0 && (val + 1) % 8 != 0) {
                                if (document.getElementById("b" + bottomright).classList[2] == "none") {
                                    document.getElementById("b" + bottomright).style.filter = "brightness(1)";
                                    arry.push(bottomright);
                                }
                                if (document.getElementById("b" + bottomright).classList[1] == "white") {
                                    document.getElementById("b" + bottomright).style.filter = "sepia(1)";
                                    arry.push(bottomright);
                                }
                            }
                            if (righttop > 0 && (righttop - 1) % 8 != 0) {
                                if (document.getElementById("b" + righttop).classList[2] == "none") {
                                    document.getElementById("b" + righttop).style.filter = "brightness(1)";
                                    arry.push(righttop);
                                }
                                if (document.getElementById("b" + righttop).classList[1] == "white") {
                                    document.getElementById("b" + righttop).style.filter = "sepia(1)";
                                    arry.push(righttop);
                                }
                            }
                            if (lefttop > 0 && lefttop % 8 != 0) {
                                if (document.getElementById("b" + lefttop).classList[2] == "none") {
                                    document.getElementById("b" + lefttop).style.filter = "brightness(1)";
                                    arry.push(lefttop);
                                }
                                if (document.getElementById("b" + lefttop).classList[1] == "white") {
                                    document.getElementById("b" + lefttop).style.filter = "sepia(1)";
                                    arry.push(lefttop);
                                }
                            }
                            if (topright > 0 && (val + 1) % 8 != 0 && val % 8 != 0) {
                                if (document.getElementById("b" + topright).classList[2] == "none") {
                                    document.getElementById("b" + topright).style.filter = "brightness(1)";
                                    arry.push(topright);
                                }
                                if (document.getElementById("b" + topright).classList[1] == "white") {
                                    document.getElementById("b" + topright).style.filter = "sepia(1)";
                                    arry.push(topright);
                                }
                            }
                            if (topleft > 0 && topleft % 8 != 0 && (topleft + 1) % 8 != 0) {
                                if (document.getElementById("b" + topleft).classList[2] == "none") {
                                    document.getElementById("b" + topleft).style.filter = "brightness(1)";
                                    arry.push(topleft);
                                }
                                if (document.getElementById("b" + topleft).classList[1] == "white") {
                                    document.getElementById("b" + topleft).style.filter = "sepia(1)";
                                    arry.push(topleft);
                                }
                            }
                        } //end of horse 
                        //black biship
                        if (pname == "biship" || pname == "queen") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            document.getElementById("b" + val).style.filter = "brightness(1)";
                            //moving bottom right
                            var bottomright = 0;
                            if (val % 8 == 0) {
                                bottomright = 0;
                            } else {
                                bottomright = 8 - val % 8;
                            }
                            j = val + 9;
                            for (var i = 0; i < bottomright; i++) {
                                if (j < 65) {
                                    if (document.getElementById("b" + j).classList[1] == "white") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j += 9;
                                    }
                                }
                            }
                            //bottom left
                            var bottomleft = 0;
                            if ((val - 1) % 8 == 0) {
                                bottomleft = 0;
                            } else if (val % 8 == 0) {
                                bottomleft = 8
                            } else {
                                bottomleft = val % 8;
                            }
                            j = val + 7;
                            for (var i = 1; i < bottomleft; i++) {
                                if (j < 65) {
                                    if (document.getElementById("b" + j).classList[1] == "white") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j += 7;
                                    }
                                }
                            }
                            //top right
                            var topright = 0;
                            if (val % 8 == 0) {
                                topright = 0;
                            } else {
                                topright = 8 - (val % 8);
                            }
                            j = val - 7;
                            for (var i = 0; i < topright; i++) {
                                if (j > 0) {
                                    if (document.getElementById("b" + j).classList[1] == "white") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j -= 7;
                                    }
                                }
                            }
                            // top left
                            var topleft = 0;
                            if ((val - 1) % 8 == 0) {
                                topleft = 0;
                            } else if (val % 8 == 0) {
                                topleft = 8
                            } else {
                                topleft = val % 8;
                            }
                            j = val - 9;
                            for (var i = 1; i < topleft; i++) {
                                if (j > 0) {
                                    if (document.getElementById("b" + j).classList[1] == "white") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j -= 9;
                                    }
                                }
                            }
                        } //end of biship
                        //black king
                        if (pname == "king") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            //right
                            if (val % 8 != 0) {
                                //right right
                                if (val + 1 < 65 && document.getElementById("b" + (val + 1)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 1)).style.filter = "brightness(1)";
                                    arry.push(val + 1);
                                }
                                if (val + 1 < 65 && document.getElementById("b" + (val + 1)).classList[1] == "white") {
                                    document.getElementById("b" + (val + 1)).style.filter = "sepia(1)";
                                    arry.push(val + 1);
                                }
                                if (val + 9 < 65 && document.getElementById("b" + (val + 9)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 9)).style.filter = "brightness(1)";
                                    arry.push(val + 9);
                                }
                                if (val + 9 < 65 && document.getElementById("b" + (val + 9)).classList[1] == "white") {
                                    document.getElementById("b" + (val + 9)).style.filter = "sepia(1)";
                                    arry.push(val + 9);
                                }
                                if (val - 7 > 0 && document.getElementById("b" + (val - 7)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 7)).style.filter = "brightness(1)";
                                    arry.push(val - 7);
                                }
                                if (val - 7 > 0 && document.getElementById("b" + (val - 7)).classList[1] == "white") {
                                    document.getElementById("b" + (val - 7)).style.filter = "sepia(1)";
                                    arry.push(val - 7);
                                }
                            }
                            //left
                            if ((val - 1) % 8 != 0) {
                                if (val - 1 > 0 && document.getElementById("b" + (val - 1)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 1)).style.filter = "brightness(1)";
                                    arry.push(val - 1);
                                }
                                if (val - 1 > 0 && document.getElementById("b" + (val - 1)).classList[1] == "white") {
                                    document.getElementById("b" + (val - 1)).style.filter = "sepia(1)";
                                    arry.push(val - 1);
                                }
                                if (val - 9 > 0 && document.getElementById("b" + (val - 9)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 9)).style.filter = "brightness(1)";
                                    arry.push(val - 9);
                                }
                                if (val - 9 > 0 && document.getElementById("b" + (val - 9)).classList[1] == "white") {
                                    document.getElementById("b" + (val - 9)).style.filter = "sepia(1)";
                                    arry.push(val - 9);
                                }
                                if (val + 7 < 65 && document.getElementById("b" + (val + 7)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 7)).style.filter = "brightness(1)";
                                    arry.push(val + 7);
                                }
                                if (val + 7 < 65 && document.getElementById("b" + (val + 7)).classList[1] == "white") {
                                    document.getElementById("b" + (val + 7)).style.filter = "sepia(1)";
                                    arry.push(val + 7);
                                }
                            }
                            //bottom
                            if (val + 8 < 65) {
                                if (document.getElementById("b" + (val + 8)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 8)).style.filter = "brightness(1)";
                                    arry.push(val + 8);
                                }
                                if (document.getElementById("b" + (val + 8)).classList[1] == "white") {
                                    document.getElementById("b" + (val + 8)).style.filter = "sepia(1)";
                                    arry.push(val + 8);
                                }
                            }
                            //top
                            if (val - 8 > 0) {
                                if (document.getElementById("b" + (val - 8)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 8)).style.filter = "brightness(1)";
                                    arry.push(val - 8);
                                }
                                if (document.getElementById("b" + (val - 8)).classList[1] == "white") {
                                    document.getElementById("b" + (val - 8)).style.filter = "sepia(1)";
                                    arry.push(val - 8);
                                }
                            }
                        }
                    } //end of black pice check
                    //checking white pice
                    if (pcolor == "white" && turn[turnval] == pturn) {
                        arry = [];
                        //white pawn move
                        if (pname == "pawn") {
                            lastp = p;
                            //for starting position
                            if (val < 57 && val > 48) {
                                //looping to get all three cells at once
                                for (var i = 0; i < 2; i++) {
                                    j -= 8;
                                    document.getElementById("b" + val).style.filter = "brightness(1)";
                                    //checking next cell is empty
                                    if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                        document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                        arry.push(val + j);
                                    }
                                } //end of for loop
                            } else {
                                j -= 8
                                document.getElementById("b" + val).style.filter = "brightness(1)";
                                //checking next cell
                                if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                }
                            }
                            //checking none last pawn
                            if ((val - 1) % 8 != 0 && document.getElementById("b" + (val - 9)).classList[1] == "black") {
                                document.getElementById("b" + (val - 9)).style.filter = "sepia(1)";
                                arry.push(val - 9);
                            }
                            //checking none fist pawn
                            if (val % 8 != 0 && document.getElementById("b" + (val - 7)).classList[1] == "black") {
                                document.getElementById("b" + (val - 7)).style.filter = "sepia(1)";
                                arry.push(val - 7);
                            }
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                        } //end of pawn if
                        //white rook move
                        if (pname == "rook" || pname == "queen") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            //bottom move
                            j = 8;
                            var bottommove = 0;
                            var test = val;
                            for (var i = 0; i < 8; i++) {
                                test += 8
                                if (test < 65) {
                                    bottommove += 1;
                                }
                            }
                            for (var i = 0; i < bottommove; i++) {
                                document.getElementById("b" + val).style.filter = "brightness(1)";
                                if (document.getElementById("b" + (val + j)).classList[1] == "black") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                //checking for empty cell
                                if (document.getElementById("b" + (val + j)).classList[2] == "none" && (val + j) < 65) {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j += 8;
                                }
                            } //end of bottom move
                            //top move
                            j = -8;
                            var topmove = 0;
                            test = val
                            for (var i = 0; i < 8; i++) {
                                test -= 8
                                if (test > 0) {
                                    topmove += 1;
                                }
                            }
                            for (var i = 0; i < topmove; i++) {
                                if (document.getElementById("b" + (val + j)).classList[1] == "black") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                //checking for empty cells
                                if ((val + j) > 0 && document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j -= 8;
                                }
                            } //end of top move
                            //right move
                            j = 1;
                            var moveright = 0;
                            //checking not at right edge
                            if (val % 8 == 0) {
                                moveright = 0
                            } else {
                                moveright = 8 - (val % 8);
                            }
                            for (var i = 0; i < moveright; i++) {
                                if (document.getElementById("b" + (val + j)).classList[1] == "black") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j += 1;
                                }
                            } //end of right move
                            //left move{
                            j = -1;
                            //checking not at left edge
                            var moveleft = 0;
                            if (val % 8 == 0) {
                                moveleft = 7;
                            } else {
                                moveleft = (val - 1) % 8;
                            }
                            for (var i = 0; i < moveleft; i++) {
                                if (document.getElementById("b" + (val + j)).classList[1] == "black") {
                                    document.getElementById("b" + (val + j)).style.filter = "sepia(1)";
                                    arry.push(val + j);
                                }
                                if (document.getElementById("b" + (val + j)).classList[2] == "none") {
                                    document.getElementById("b" + (val + j)).style.filter = "brightness(1)";
                                    arry.push(val + j);
                                    j -= 1;
                                }
                            }
                        } //end of rook
                        //white horse
                        if (pname == "horse") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            var leftbottom = val + 8 + 7;
                            var rightbottom = val + 8 + 9;
                            var bottomright = val + 10;
                            var bottomleft = val + 6;
                            var righttop = val - 8 - 7;
                            var lefttop = val - 8 - 9;
                            var topright = val - 6;
                            var topleft = val - 10;
                            document.getElementById("b" + val).style.filter = "brightness(1)";
                            if (leftbottom < 65 && leftbottom % 8 != 0) {
                                if (document.getElementById("b" + leftbottom).classList[2] == "none") {
                                    document.getElementById("b" + leftbottom).style.filter = "brightness(1)";
                                    arry.push(leftbottom);
                                }
                                if (document.getElementById("b" + leftbottom).classList[1] == "black") {
                                    document.getElementById("b" + leftbottom).style.filter = "sepia(1)";
                                    arry.push(leftbottom);
                                }
                            }
                            if (rightbottom < 65 && val % 8 != 0) {
                                if (document.getElementById("b" + rightbottom).classList[2] == "none") {
                                    document.getElementById("b" + rightbottom).style.filter = "brightness(1)";
                                    arry.push(rightbottom);
                                }
                                if (document.getElementById("b" + rightbottom).classList[1] == "black") {
                                    document.getElementById("b" + rightbottom).style.filter = "sepia(1)";
                                    arry.push(rightbottom);
                                }
                            }
                            if (bottomleft < 65 && (val - 1) % 8 != 0 && (val - 2) % 8 != 0) {
                                if (document.getElementById("b" + bottomleft).classList[2] == "none") {
                                    document.getElementById("b" + bottomleft).style.filter = "brightness(1)";
                                    arry.push(bottomleft);
                                }
                                if (document.getElementById("b" + bottomleft).classList[1] == "black") {
                                    document.getElementById("b" + bottomleft).style.filter = "sepia(1)";
                                    arry.push(bottomleft);
                                }
                            }
                            if (bottomright < 65 && val % 8 != 0 && (val + 1) % 8 != 0) {
                                if (document.getElementById("b" + bottomright).classList[2] == "none") {
                                    document.getElementById("b" + bottomright).style.filter = "brightness(1)";
                                    arry.push(bottomright);
                                }
                                if (document.getElementById("b" + bottomright).classList[1] == "white") {
                                    document.getElementById("b" + bottomright).style.filter = "sepia(1)";
                                    arry.push(bottomright);
                                }
                            }
                            if (righttop > 0 && (righttop - 1) % 8 != 0) {
                                if (document.getElementById("b" + righttop).classList[2] == "none") {
                                    document.getElementById("b" + righttop).style.filter = "brightness(1)";
                                    arry.push(righttop);
                                }
                                if (document.getElementById("b" + righttop).classList[1] == "black") {
                                    document.getElementById("b" + righttop).style.filter = "sepia(1)";
                                    arry.push(righttop);
                                }
                            }
                            if (lefttop > 0 && lefttop % 8 != 0) {
                                if (document.getElementById("b" + lefttop).classList[2] == "none") {
                                    document.getElementById("b" + lefttop).style.filter = "brightness(1)";
                                    arry.push(lefttop);
                                }
                                if (document.getElementById("b" + lefttop).classList[1] == "black") {
                                    document.getElementById("b" + lefttop).style.filter = "sepia(1)";
                                    arry.push(lefttop);
                                }
                            }
                            if (topright > 0 && (val + 1) % 8 != 0 && val % 8 != 0) {
                                if (document.getElementById("b" + topright).classList[2] == "none") {
                                    document.getElementById("b" + topright).style.filter = "brightness(1)";
                                    arry.push(topright);
                                }
                                if (document.getElementById("b" + topright).classList[1] == "black") {
                                    document.getElementById("b" + topright).style.filter = "sepia(1)";
                                    arry.push(topright);
                                }
                            }
                            if (topleft > 0 && topleft % 8 != 0 && (topleft + 1) % 8 != 0) {
                                if (document.getElementById("b" + topleft).classList[2] == "none") {
                                    document.getElementById("b" + topleft).style.filter = "brightness(1)";
                                    arry.push(topleft);
                                }
                                if (document.getElementById("b" + topleft).classList[1] == "black") {
                                    document.getElementById("b" + topleft).style.filter = "sepia(1)";
                                    arry.push(topleft);
                                }
                            }
                        } //end of horse 
                        //white biship
                        if (pname == "biship" || pname == "queen") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            document.getElementById("b" + val).style.filter = "brightness(1)";
                            //moving bottom right
                            var bottomright = 0;
                            if (val % 8 == 0) {
                                bottomright = 0;
                            } else {
                                bottomright = 8 - val % 8;
                            }
                            j = val + 9;
                            for (var i = 0; i < bottomright; i++) {
                                if (j < 65) {
                                    if (document.getElementById("b" + j).classList[1] == "black") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j += 9;
                                    }
                                }
                            }
                            //bottom left
                            var bottomleft = 0;
                            if ((val - 1) % 8 == 0) {
                                bottomleft = 0;
                            } else if (val % 8 == 0) {
                                bottomleft = 8
                            } else {
                                bottomleft = val % 8;
                            }
                            j = val + 7;
                            for (var i = 1; i < bottomleft; i++) {
                                if (j < 65) {
                                    if (document.getElementById("b" + j).classList[1] == "black") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j += 7;
                                    }
                                }
                            }
                            //top right
                            var topright = 0;
                            if (val % 8 == 0) {
                                topright = 0;
                            } else {
                                topright = 8 - (val % 8);
                            }
                            j = val - 7;
                            for (var i = 0; i < topright; i++) {
                                if (j > 0) {
                                    if (document.getElementById("b" + j).classList[1] == "black") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j -= 7;
                                    }
                                }
                            }
                            // top left
                            var topleft = 0;
                            if ((val - 1) % 8 == 0) {
                                topleft = 0;
                            } else if (val % 8 == 0) {
                                topleft = 8
                            } else {
                                topleft = val % 8;
                            }
                            j = val - 9;
                            for (var i = 1; i < topleft; i++) {
                                if (j > 0) {
                                    if (document.getElementById("b" + j).classList[1] == "black") {
                                        document.getElementById("b" + j).style.filter = "sepia(1)";
                                        arry.push(j);
                                    }
                                    if (document.getElementById("b" + j).classList[2] == "none") {
                                        document.getElementById("b" + j).style.filter = "brightness(1)";
                                        arry.push(j);
                                        j -= 9;
                                    }
                                }
                            }
                        } //end of biship
                        //white king
                        if (pname == "king") {
                            lastp = p;
                            lpcolor = pcolor;
                            lpname = pname;
                            lpturn = pturn;
                            lastval = val;
                            //right
                            if (val % 8 != 0) {
                                //right right
                                if (val + 1 < 65 && document.getElementById("b" + (val + 1)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 1)).style.filter = "brightness(1)";
                                    arry.push(val + 1);
                                }
                                if (val + 1 < 65 && document.getElementById("b" + (val + 1)).classList[1] == "black") {
                                    document.getElementById("b" + (val + 1)).style.filter = "sepia(1)";
                                    arry.push(val + 1);
                                }
                                if (val + 9 < 65 && document.getElementById("b" + (val + 9)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 9)).style.filter = "brightness(1)";
                                    arry.push(val + 9);
                                }
                                if (val + 9 < 65 && document.getElementById("b" + (val + 9)).classList[1] == "black") {
                                    document.getElementById("b" + (val + 9)).style.filter = "sepia(1)";
                                    arry.push(val + 9);
                                }
                                if (val - 7 > 0 && document.getElementById("b" + (val - 7)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 7)).style.filter = "brightness(1)";
                                    arry.push(val - 7);
                                }
                                if (val - 7 > 0 && document.getElementById("b" + (val - 7)).classList[1] == "black") {
                                    document.getElementById("b" + (val - 7)).style.filter = "sepia(1)";
                                    arry.push(val - 7);
                                }
                            }
                            //left
                            if ((val - 1) % 8 != 0) {
                                if (val - 1 > 0 && document.getElementById("b" + (val - 1)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 1)).style.filter = "brightness(1)";
                                    arry.push(val - 1);
                                }
                                if (val - 1 > 0 && document.getElementById("b" + (val - 1)).classList[1] == "black") {
                                    document.getElementById("b" + (val - 1)).style.filter = "sepia(1)";
                                    arry.push(val - 1);
                                }
                                if (val - 9 > 0 && document.getElementById("b" + (val - 9)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 9)).style.filter = "brightness(1)";
                                    arry.push(val - 9);
                                }
                                if (val - 9 > 0 && document.getElementById("b" + (val - 9)).classList[1] == "black") {
                                    document.getElementById("b" + (val - 9)).style.filter = "sepia(1)";
                                    arry.push(val - 9);
                                }
                                if (val + 7 < 65 && document.getElementById("b" + (val + 7)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 7)).style.filter = "brightness(1)";
                                    arry.push(val + 7);
                                }
                                if (val + 7 < 65 && document.getElementById("b" + (val + 7)).classList[1] == "black") {
                                    document.getElementById("b" + (val + 7)).style.filter = "sepia(1)";
                                    arry.push(val + 7);
                                }
                            }
                            //bottom
                            if (val + 8 < 65) {
                                if (document.getElementById("b" + (val + 8)).classList[2] == "none") {
                                    document.getElementById("b" + (val + 8)).style.filter = "brightness(1)";
                                    arry.push(val + 8);
                                }
                                if (document.getElementById("b" + (val + 8)).classList[1] == "black") {
                                    document.getElementById("b" + (val + 8)).style.filter = "sepia(1)";
                                    arry.push(val + 8);
                                }
                            }
                            //top
                            if (val - 8 > 0) {
                                if (document.getElementById("b" + (val - 8)).classList[2] == "none") {
                                    document.getElementById("b" + (val - 8)).style.filter = "brightness(1)";
                                    arry.push(val - 8);
                                }
                                if (document.getElementById("b" + (val - 8)).classList[1] == "black") {
                                    document.getElementById("b" + (val - 8)).style.filter = "sepia(1)";
                                    arry.push(val - 8);
                                }
                            }
                        }
                    } //end of white pice check
                } //end of cheching empty cells
                move(val);
            }
            function move(val) {
                for (i = 0; i < arry.length; i++) {
                    if (val == arry[i]) {
                        //black pawn queen
                        if (arry[i] > 56 && arry[i] <= 64 && lpname == "pawn") {
                            lpname = "queen"
                            lastp = "<p class='p'>&#" + bq + ";</p>";
                        }
                        //white pawn queen
                        if (arry[i] > 0 && arry[i] <= 8 && lpname == "pawn") {
                            lpname = "queen"
                            lastp = "<p class='p'>&#" + wq + ";</p>";
                        }
                        cls = "box " + lpcolor + " " + lpname + " " + lpturn;
                        id.innerHTML = lastp;
                        id.className = cls;
                        document.getElementById("b" + lastval).innerHTML = "";
                        document.getElementById("b" + lastval).classList = "box pcolor none nturn";
                        arry = [];
                        if (turnval == 1) {
                            turnval = 0;
                        } else {
                            turnval = 1;
                        }
                    }
                }
                check();
            }
            function check() {
                var bpawn = document.getElementsByClassName("black pawn").length;
                var brook = document.getElementsByClassName("black rook").length;
                var bhorse = document.getElementsByClassName("black horse").length;
                var bbiship = document.getElementsByClassName("black biship").length;
                var bqueen = document.getElementsByClassName("black queen").length;
                var bking = document.getElementsByClassName("black king").length;
                var wpawn = document.getElementsByClassName("white pawn").length;
                var wrook = document.getElementsByClassName("white rook").length;
                var whorse = document.getElementsByClassName("white horse").length;
                var wbiship = document.getElementsByClassName("white biship").length;
                var wqueen = document.getElementsByClassName("white queen").length;
                var wking = document.getElementsByClassName("white king").length;
                if (bpawn < 8) {
                    html = "";
                    for (var i = 0; i < (8 - bpawn); i++) {
                        html += "&#" + bp + ";"
                    }
                    document.getElementById("bpawn").innerHTML = html;
                }
                if (brook < 2) {
                    html = "";
                    for (var i = 0; i < (2 - brook); i++) {
                        html += "&#" + br + ";"
                    }
                    document.getElementById("brook").innerHTML = html;
                }
                if (bhorse < 2) {
                    html = "";
                    for (var i = 0; i < (2 - bhorse); i++) {
                        html += "&#" + bh + ";"
                    }
                    document.getElementById("bhorse").innerHTML = html;
                }
                if (bbiship < 2) {
                    html = "";
                    for (var i = 0; i < (2 - bbiship); i++) {
                        html += "&#" + bb + ";"
                    }
                    document.getElementById("bbiship").innerHTML = html;
                }
                if (bqueen < 1) {
                    html = "";
                    for (var i = 0; i < (1 - bqueen); i++) {
                        html += "&#" + bq + ";"
                    }
                    document.getElementById("bqueen").innerHTML = html;
                }
                if (wpawn < 8) {
                    html = "";
                    for (var i = 0; i < (8 - wpawn); i++) {
                        html += "&#" + wp + ";"
                    }
                    document.getElementById("wpawn").innerHTML = html;
                }
                if (wrook < 2) {
                    html = "";
                    for (var i = 0; i < (2 - wrook); i++) {
                        html += "&#" + wr + ";"
                    }
                    document.getElementById("wrook").innerHTML = html;
                }
                if (whorse < 2) {
                    html = "";
                    for (var i = 0; i < (2 - whorse); i++) {
                        html += "&#" + wh + ";"
                    }
                    document.getElementById("whorse").innerHTML = html;
                }
                if (wbiship < 2) {
                    html = "";
                    for (var i = 0; i < (2 - wbiship); i++) {
                        html += "&#" + wb + ";"
                    }
                    document.getElementById("wbiship").innerHTML = html;
                }
                if (wqueen < 1) {
                    html = "";
                    for (var i = 0; i < (1 - wqueen); i++) {
                        html += "&#" + wq + ";"
                    }
                    document.getElementById("wqueen").innerHTML = html;
                }
                if (bking < 1) {
                    document.getElementById("winner").style.display = "flex";
                    document.getElementById("winner").innerHTML = "<p>White has won the game</p>";
                }
                if (wking < 1) {
                    document.getElementById("winner").style.display = "flex";
                    document.getElementById("winner").innerHTML = "<p>Black has won the game</p>";
                }
                if (turnval == 1) {
                    for (var i = 0; i < document.getElementsByClassName("box").length; i++) {
                        document.getElementsByClassName("box")[i].style.boxShadow = "0 0 0px 0px red";
                    }
                    for (var i = 0; i < document.getElementsByClassName("black").length; i++) {
                        document.getElementsByClassName("black")[i].style.boxShadow = "0 0 2px 1px red";
                    }
                } else {
                    for (var i = 0; i < document.getElementsByClassName("box").length; i++) {
                        document.getElementsByClassName("box")[i].style.boxShadow = "0 0 0px 0px red";
                    }
                    for (var i = 0; i < document.getElementsByClassName("white").length; i++) {
                        document.getElementsByClassName("white")[i].style.boxShadow = "0 0 2px 1px red";
                    }
                }
            }
</script>
        <div id="blackkill">
            <p id="bpawn"></p>
            <p id="brook"></p>
            <p id="bhorse"></p>
            <p id="bbiship"></p>
            <p id="bqueen"></p>
            <p id="bking"></p>
        </div>
        <div id="board"></div>
        <div id="whitekill">
            <p id="wpawn"></p>
            <p id="wrook"></p>
            <p id="whorse"></p>
            <p id="wbiship"></p>
            <p id="wqueen"></p>
            <p id="wking"></p>
        </div>
        <div id="winner"></div>
    </body>

    </html>
    `))
	})

	r.GET("/applications", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
        <!DOCTYPE html> <html> <head> <meta name="viewport" content="width=device-width, initial-scale=1"> <style> body, html { height: 100%; margin: 0; } .content { position: absolute; top: 15%; left:25%; background: rgb(0, 0, 0); /* Fallback color */ background: rgba(0, 0, 0, 0.76); /* Black background with 0.5 opacity */ color: #f1f1f1; width: 50%; padding: 20px; } .bg { /* The image used */ background-image: url("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRcRqNquWxQHJAPgugDwzXokAU_dQUXzknUTA&usqp=CAU"); /* Full height */ height: 100%; /* Center and scale the image nicely */ background-position: center; background-repeat: no-repeat; background-size: cover; } table { font-family: arial, sans-serif; border-collapse: collapse; width: 100%; } td, th { border: 1px solid #dddddd; text-align: left; padding: 8px; } div.parent { text-align: center; } ul { display: inline-block; text-align: left; }</style>
    </head> <body> <div class="bg"></div>  <div class="content"> <h1 id="home" 
        style="text-align: center;font-weight: bold;text-decoration: underline;">
            WELCOME TO JAVASCRIPT APPLICATIONS!!</h1> 
   <h3 style="text-align:center;"> Click on any of the below JavaScript apps!</h3>
   <div class="parent"> <ul>
        <li><a href='/api/calculator'>Calculator</a></li> 
       <li><a href='/api/maze'>Maze</a></li>
        <li><a href='/api/tic_tac_toe'>Tic-tac-toe</a></li>
        <li><a href='/api/clock'>Analogue clock</a></li>
        <li><a href='/api/hangman'>Hangman</a></li>
        <li><a href='/api/puzzles'>Estonian puzzles</a></li>
        <li><a href='/api/sudoku'>Sudoku</a></li>
        <li><a href='/api/virtual_keyboard'>Virtual keyboard</a></li>
	<li><a href='/api/solitaire'>Solitaire</a></li>
	<li><a href='/api/chess'>Chess</a></li>
	<li><a href='/api/dino'>Dino</a></li>
	<li><a href='/api/sass_'>Sass</a></li>
	<li><a href='/api/tilt_maze'>Tilted Maze</a></li>
	<li><a href='/api/codepen'>Codepen</a></li>
   </ul> </div> 
       </div> </body> </html>
    `))
	})

	r.GET("/calculator", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
        <html>
    <head>
    <style>
    html {
      font-size: 62.5%;
      box-sizing: border-box;
    }
    
    *, *::before, *::after {
      margin: 0;
      padding: 0;
      box-sizing: inherit;
    }
    
    .calculator {
      border: 1px solid #ccc;
      border-radius: 5px;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 400px;
    }
    
    .calculator-screen {
      width: 100%;
      font-size: 5rem;
      height: 80px;
      border: none;
      background-color: #252525;
      color: #fff;
      text-align: right;
      padding-right: 20px;
      padding-left: 10px;
    }
    
    button {
      height: 60px;
      background-color: #fff;
      border-radius: 3px;
      border: 1px solid #c4c4c4;
      background-color: transparent;
      font-size: 2rem;
      color: #333;
      background-image: linear-gradient(to bottom,transparent,transparent 50%,rgba(0,0,0,.04));
      box-shadow: inset 0 0 0 1px rgba(255,255,255,.05), inset 0 1px 0 0 rgba(255,255,255,.45), inset 0 -1px 0 0 rgba(255,255,255,.15), 0 1px 0 0 rgba(255,255,255,.15);
      text-shadow: 0 1px rgba(255,255,255,.4);
    }
    
    button:hover {
      background-color: #eaeaea;
    }
    
    .operator {
      color: #337cac;
    }
    
    .all-clear {
      background-color: #f0595f;
      border-color: #b0353a;
      color: #fff;
    }
    
    .all-clear:hover {
      background-color: #f17377;
    }
    
    .equal-sign {
      background-color: #2e86c0;
      border-color: #337cac;
      color: #fff;
      height: 100%;
      grid-area: 2 / 4 / 6 / 5;
    }
    
    .equal-sign:hover {
      background-color: #4e9ed4;
    }
    
    .calculator-keys {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      grid-gap: 20px;
      padding: 20px;
    }
    </style>
    </head>
    <body>
    <div class="calculator">
    
      <input type="text" class="calculator-screen" value="" disabled />
      
      <div class="calculator-keys">
        
        <button type="button" class="operator" value="+">+</button>
        <button type="button" class="operator" value="-">-</button>
        <button type="button" class="operator" value="*">&times;</button>
        <button type="button" class="operator" value="/">&divide;</button>
    
        <button type="button" value="7">7</button>
        <button type="button" value="8">8</button>
        <button type="button" value="9">9</button>
    
    
        <button type="button" value="4">4</button>
        <button type="button" value="5">5</button>
        <button type="button" value="6">6</button>
    
    
        <button type="button" value="1">1</button>
        <button type="button" value="2">2</button>
        <button type="button" value="3">3</button>
    
    
        <button type="button" value="0">0</button>
        <button type="button" class="decimal" value=".">.</button>
        <button type="button" class="all-clear" value="all-clear">AC</button>
    
        <button type="button" class="equal-sign operator" value="=">=</button>
    
      </div>
    </div>
    <script>
    const calculator = {
      displayValue: '0',
      firstOperand: null,
      waitingForSecondOperand: false,
      operator: null,
    };
    
    function inputDigit(digit) {
      const { displayValue, waitingForSecondOperand } = calculator;
    
      if (waitingForSecondOperand === true) {
        calculator.displayValue = digit;
        calculator.waitingForSecondOperand = false;
      } else {
        calculator.displayValue = displayValue === '0' ? digit : displayValue + digit;
      }
    }
    
    function inputDecimal(dot) {
      if (calculator.waitingForSecondOperand === true) {
          calculator.displayValue = "0."
        calculator.waitingForSecondOperand = false;
        return
      }
    
      if (!calculator.displayValue.includes(dot)) {
        calculator.displayValue += dot;
      }
    }
    
    function handleOperator(nextOperator) {
      const { firstOperand, displayValue, operator } = calculator
      const inputValue = parseFloat(displayValue);
      
      if (operator && calculator.waitingForSecondOperand)  {
        calculator.operator = nextOperator;
        return;
      }
    
    
      if (firstOperand == null && !isNaN(inputValue)) {
        calculator.firstOperand = inputValue;
      } else if (operator) {
        const result = calculate(firstOperand, inputValue, operator);
    
        calculator.displayValue = `+"`"+`\${parseFloat(result.toFixed(7))}`+"`"+`;
        calculator.firstOperand = result;
      }
    
      calculator.waitingForSecondOperand = true;
      calculator.operator = nextOperator;
    }
    
    function calculate(firstOperand, secondOperand, operator) {
      if (operator === '+') {
        return firstOperand + secondOperand;
      } else if (operator === '-') {
        return firstOperand - secondOperand;
      } else if (operator === '*') {
        return firstOperand * secondOperand;
      } else if (operator === '/') {
        return firstOperand / secondOperand;
      }
    
      return secondOperand;
    }
    
    function resetCalculator() {
      calculator.displayValue = '0';
      calculator.firstOperand = null;
      calculator.waitingForSecondOperand = false;
      calculator.operator = null;
    }
    
    function updateDisplay() {
      const display = document.querySelector('.calculator-screen');
      display.value = calculator.displayValue;
    }
    
    updateDisplay();
    
    const keys = document.querySelector('.calculator-keys');
    keys.addEventListener('click', event => {
      const { target } = event;
      const { value } = target;
      if (!target.matches('button')) {
        return;
      }
    
      switch (value) {
        case '+':
        case '-':
        case '*':
        case '/':
        case '=':
          handleOperator(value);
          break;
        case '.':
          inputDecimal(value);
          break;
        case 'all-clear':
          resetCalculator();
          break;
        default:
          if (Number.isInteger(parseFloat(value))) {
            inputDigit(value);
          }
      }
    
      updateDisplay();
    });
    </script>
    </body>
    </html>
    `))
	})

	r.GET("/maze", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
<html lang="en-GB">
    <head>
      <meta charset="utf-8">
      <style>
        $menuHeight: 65px+10px;
    @mixin transition {
        transition-property: background-color opacity;
        transition-duration: 0.2s;
        transition-timing-function: ease-in-out;
    }
    
    html,
    body {
        width: 100vw;
        height: 100vh;
        position: fixed;
        padding: 0;
        margin: 0;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
    }
    
    #mazeContainer {
        transition-property: opacity;
        transition-duration: 1s;
        transition-timing-function: linear;
        top: $menuHeight;
        opacity: 0;
        display: inline-block;
        background-color: rgba(0, 0, 0, 0.30);
        margin: auto;
    
        #mazeCanvas {
            margin: 0;
            display: block;
            border: solid 1px black;
        }
    }
    
    input,
    select {
        @include transition;
        cursor: pointer;
        background-color: rgba(0, 0, 0, 0.30);
        height: 45px;
        width: 150px;
        padding: 10px;
        border: none;
        border-radius: 5px;
        color: white;
        display: inline-block;
        font-size: 15px;
        text-align: center;
        text-decoration: none;
        appearance: none;
        &:hover {
            background-color: rgba(0, 0, 0, 0.70);
        }
        &:active {
            background-color: black;
        }
        &:focus {
            outline: none;
        }
    }
    
    
    .custom-select {
        display: inline-block;
        select {
            -webkit-appearance: none;
            -moz-appearance: none;
            appearance: none;
            background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAh0lEQVQ4T93TMQrCUAzG8V9x8QziiYSuXdzFC7h4AcELOPQAdXYovZCHEATlgQV5GFTe1ozJlz/kS1IpjKqw3wQBVyy++JI0y1GTe7DCBbMAckeNIQKk/BanALBB+16LtnDELoMcsM/BESDlz2heDR3WePwKSLo5eoxz3z6NNcFD+vu3ij14Aqz/DxGbKB7CAAAAAElFTkSuQmCC');
            background-repeat: no-repeat;
            background-position: 125px center;
        }
    }
    
    #Message-Container {
        visibility: hidden;
        color: white;
        display: block;
        width: 100vw;
        height: 100vh;
        position: fixed;
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        background-color: rgba(0, 0, 0, 0.30);
        z-index: 1;
        #message {
            width: 300px;
            height: 300px;
            position: fixed;
            top: 50%;
            left: 50%;
            margin-left: -150px;
            margin-top: -150px;
        }
    }
    
    #page {
        font-family: "Segoe UI", Arial, sans-serif;
        text-align: center;
        height: auto;
        width: auto;
        margin: auto;
        #menu {
            margin: auto;
            padding: 10px;
            height: 65px;
            box-sizing: border-box;
            h1 {
                margin: 0;
                margin-bottom: 10px;
                font-weight: 600;
                font-size: 3.2rem;
            }
        }
        #view {
            position: absolute;
            top:65px;
            bottom: 0;
            left: 0;
            right: 0;
            width: 100%;
            height: auto;
               
        }
    }
    
    .border {
        border: 1px black solid;
        border-radius: 5px;
    }
    
    
    
    #gradient {
        z-index: -1;
        position: fixed;
        top: 0;
        bottom: 0;
        width: 100vw;
        height: 100vh;
        color: #fff;
        background: linear-gradient(-45deg, #EE7752, #E73C7E, #23A6D5, #23D5AB);
        background-size: 400% 400%;
        animation: Gradient 15s ease infinite;
    }
    
    @keyframes Gradient {
        0% {
            background-position: 0% 50%
        }
        50% {
            background-position: 100% 50%
        }
        100% {
            background-position: 0% 50%
        }
    }
    
     /* Extra small devices (phones, 600px and down) */
     @media only screen and (max-width: 400px) {
         input, select{
             width: 120px;
         }
     }
    
      </style>
      <body>
        <div id="gradient"></div>
        <div id="page">
          <div id="Message-Container">
            <div id="message">
              <h1>Congratulations!</h1>
              <p>You are done.</p>
              <p id="moves"></p>
              <input id="okBtn" type="button" onclick="toggleVisablity('Message-Container')" value="Cool!" />
            </div>
          </div>
          <div id="menu">
            <div class="custom-select">
              <select id="diffSelect">
                        <option value="10">Easy</option>
                        <option value="15">Medium</option>
                        <option value="25">Hard</option>
                        <option value="38">Extreme</option>                                      
                    </select>
            </div>
            <input id="startMazeBtn" type="button" onclick="makeMaze()" value="Start" />
          </div>
          <div id="view">
            <div id="mazeContainer">
              <canvas id="mazeCanvas" class="border" height="1100" width="1100"></canvas>
            </div>
          </div>
        </div>
        <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
        <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery.touchswipe/1.6.18/jquery.touchSwipe.min.js"></script>
        <script>
          function rand(max) {
      return Math.floor(Math.random() * max);
    }
    
    function shuffle(a) {
      for (let i = a.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [a[i], a[j]] = [a[j], a[i]];
      }
      return a;
    }
    
    function changeBrightness(factor, sprite) {
      var virtCanvas = document.createElement("canvas");
      virtCanvas.width = 500;
      virtCanvas.height = 500;
      var context = virtCanvas.getContext("2d");
      context.drawImage(sprite, 0, 0, 500, 500);
    
      var imgData = context.getImageData(0, 0, 500, 500);
    
      for (let i = 0; i < imgData.data.length; i += 4) {
        imgData.data[i] = imgData.data[i] * factor;
        imgData.data[i + 1] = imgData.data[i + 1] * factor;
        imgData.data[i + 2] = imgData.data[i + 2] * factor;
      }
      context.putImageData(imgData, 0, 0);
    
      var spriteOutput = new Image();
      spriteOutput.src = virtCanvas.toDataURL();
      virtCanvas.remove();
      return spriteOutput;
    }
    
    function displayVictoryMess(moves) {
      document.getElementById("moves").innerHTML = "You Moved " + moves + " Steps.";
      toggleVisablity("Message-Container");  
    }
    
    function toggleVisablity(id) {
      if (document.getElementById(id).style.visibility == "visible") {
        document.getElementById(id).style.visibility = "hidden";
      } else {
        document.getElementById(id).style.visibility = "visible";
      }
    }
    
    function Maze(Width, Height) {
      var mazeMap;
      var width = Width;
      var height = Height;
      var startCoord, endCoord;
      var dirs = ["n", "s", "e", "w"];
      var modDir = {
        n: {
          y: -1,
          x: 0,
          o: "s"
        },
        s: {
          y: 1,
          x: 0,
          o: "n"
        },
        e: {
          y: 0,
          x: 1,
          o: "w"
        },
        w: {
          y: 0,
          x: -1,
          o: "e"
        }
      };
    
      this.map = function() {
        return mazeMap;
      };
      this.startCoord = function() {
        return startCoord;
      };
      this.endCoord = function() {
        return endCoord;
      };
    
      function genMap() {
        mazeMap = new Array(height);
        for (y = 0; y < height; y++) {
          mazeMap[y] = new Array(width);
          for (x = 0; x < width; ++x) {
            mazeMap[y][x] = {
              n: false,
              s: false,
              e: false,
              w: false,
              visited: false,
              priorPos: null
            };
          }
        }
      }
    
      function defineMaze() {
        var isComp = false;
        var move = false;
        var cellsVisited = 1;
        var numLoops = 0;
        var maxLoops = 0;
        var pos = {
          x: 0,
          y: 0
        };
        var numCells = width * height;
        while (!isComp) {
          move = false;
          mazeMap[pos.x][pos.y].visited = true;
    
          if (numLoops >= maxLoops) {
            shuffle(dirs);
            maxLoops = Math.round(rand(height / 8));
            numLoops = 0;
          }
          numLoops++;
          for (index = 0; index < dirs.length; index++) {
            var direction = dirs[index];
            var nx = pos.x + modDir[direction].x;
            var ny = pos.y + modDir[direction].y;
    
            if (nx >= 0 && nx < width && ny >= 0 && ny < height) {
              //Check if the tile is already visited
              if (!mazeMap[nx][ny].visited) {
                //Carve through walls from this tile to next
                mazeMap[pos.x][pos.y][direction] = true;
                mazeMap[nx][ny][modDir[direction].o] = true;
    
                //Set Currentcell as next cells Prior visited
                mazeMap[nx][ny].priorPos = pos;
                //Update Cell position to newly visited location
                pos = {
                  x: nx,
                  y: ny
                };
                cellsVisited++;
                //Recursively call this method on the next tile
                move = true;
                break;
              }
            }
          }
    
          if (!move) {
            //  If it failed to find a direction,
            //  move the current position back to the prior cell and Recall the method.
            pos = mazeMap[pos.x][pos.y].priorPos;
          }
          if (numCells == cellsVisited) {
            isComp = true;
          }
        }
      }
    
      function defineStartEnd() {
        switch (rand(4)) {
          case 0:
            startCoord = {
              x: 0,
              y: 0
            };
            endCoord = {
              x: height - 1,
              y: width - 1
            };
            break;
          case 1:
            startCoord = {
              x: 0,
              y: width - 1
            };
            endCoord = {
              x: height - 1,
              y: 0
            };
            break;
          case 2:
            startCoord = {
              x: height - 1,
              y: 0
            };
            endCoord = {
              x: 0,
              y: width - 1
            };
            break;
          case 3:
            startCoord = {
              x: height - 1,
              y: width - 1
            };
            endCoord = {
              x: 0,
              y: 0
            };
            break;
        }
      }
    
      genMap();
      defineStartEnd();
      defineMaze();
    }
    
    function DrawMaze(Maze, ctx, cellsize, endSprite = null) {
      var map = Maze.map();
      var cellSize = cellsize;
      var drawEndMethod;
      ctx.lineWidth = cellSize / 40;
    
      this.redrawMaze = function(size) {
        cellSize = size;
        ctx.lineWidth = cellSize / 50;
        drawMap();
        drawEndMethod();
      };
    
      function drawCell(xCord, yCord, cell) {
        var x = xCord * cellSize;
        var y = yCord * cellSize;
    
        if (cell.n == false) {
          ctx.beginPath();
          ctx.moveTo(x, y);
          ctx.lineTo(x + cellSize, y);
          ctx.stroke();
        }
        if (cell.s === false) {
          ctx.beginPath();
          ctx.moveTo(x, y + cellSize);
          ctx.lineTo(x + cellSize, y + cellSize);
          ctx.stroke();
        }
        if (cell.e === false) {
          ctx.beginPath();
          ctx.moveTo(x + cellSize, y);
          ctx.lineTo(x + cellSize, y + cellSize);
          ctx.stroke();
        }
        if (cell.w === false) {
          ctx.beginPath();
          ctx.moveTo(x, y);
          ctx.lineTo(x, y + cellSize);
          ctx.stroke();
        }
      }
    
      function drawMap() {
        for (x = 0; x < map.length; x++) {
          for (y = 0; y < map[x].length; y++) {
            drawCell(x, y, map[x][y]);
          }
        }
      }
    
      function drawEndFlag() {
        var coord = Maze.endCoord();
        var gridSize = 4;
        var fraction = cellSize / gridSize - 2;
        var colorSwap = true;
        for (let y = 0; y < gridSize; y++) {
          if (gridSize % 2 == 0) {
            colorSwap = !colorSwap;
          }
          for (let x = 0; x < gridSize; x++) {
            ctx.beginPath();
            ctx.rect(
              coord.x * cellSize + x * fraction + 4.5,
              coord.y * cellSize + y * fraction + 4.5,
              fraction,
              fraction
            );
            if (colorSwap) {
              ctx.fillStyle = "rgba(0, 0, 0, 0.8)";
            } else {
              ctx.fillStyle = "rgba(255, 255, 255, 0.8)";
            }
            ctx.fill();
            colorSwap = !colorSwap;
          }
        }
      }
    
      function drawEndSprite() {
        var offsetLeft = cellSize / 50;
        var offsetRight = cellSize / 25;
        var coord = Maze.endCoord();
        ctx.drawImage(
          endSprite,
          2,
          2,
          endSprite.width,
          endSprite.height,
          coord.x * cellSize + offsetLeft,
          coord.y * cellSize + offsetLeft,
          cellSize - offsetRight,
          cellSize - offsetRight
        );
      }
    
      function clear() {
        var canvasSize = cellSize * map.length;
        ctx.clearRect(0, 0, canvasSize, canvasSize);
      }
    
      if (endSprite != null) {
        drawEndMethod = drawEndSprite;
      } else {
        drawEndMethod = drawEndFlag;
      }
      clear();
      drawMap();
      drawEndMethod();
    }
    
    function Player(maze, c, _cellsize, onComplete, sprite = null) {
      var ctx = c.getContext("2d");
      var drawSprite;
      var moves = 0;
      drawSprite = drawSpriteCircle;
      if (sprite != null) {
        drawSprite = drawSpriteImg;
      }
      var player = this;
      var map = maze.map();
      var cellCoords = {
        x: maze.startCoord().x,
        y: maze.startCoord().y
      };
      var cellSize = _cellsize;
      var halfCellSize = cellSize / 2;
    
      this.redrawPlayer = function(_cellsize) {
        cellSize = _cellsize;
        drawSpriteImg(cellCoords);
      };
    
      function drawSpriteCircle(coord) {
        ctx.beginPath();
        ctx.fillStyle = "yellow";
        ctx.arc(
          (coord.x + 1) * cellSize - halfCellSize,
          (coord.y + 1) * cellSize - halfCellSize,
          halfCellSize - 2,
          0,
          2 * Math.PI
        );
        ctx.fill();
        if (coord.x === maze.endCoord().x && coord.y === maze.endCoord().y) {
          onComplete(moves);
          player.unbindKeyDown();
        }
      }
    
      function drawSpriteImg(coord) {
        var offsetLeft = cellSize / 50;
        var offsetRight = cellSize / 25;
        ctx.drawImage(
          sprite,
          0,
          0,
          sprite.width,
          sprite.height,
          coord.x * cellSize + offsetLeft,
          coord.y * cellSize + offsetLeft,
          cellSize - offsetRight,
          cellSize - offsetRight
        );
        if (coord.x === maze.endCoord().x && coord.y === maze.endCoord().y) {
          onComplete(moves);
          player.unbindKeyDown();
        }
      }
    
      function removeSprite(coord) {
        var offsetLeft = cellSize / 50;
        var offsetRight = cellSize / 25;
        ctx.clearRect(
          coord.x * cellSize + offsetLeft,
          coord.y * cellSize + offsetLeft,
          cellSize - offsetRight,
          cellSize - offsetRight
        );
      }
    
      function check(e) {
        var cell = map[cellCoords.x][cellCoords.y];
        moves++;
        switch (e.keyCode) {
          case 65:
          case 37: // west
            if (cell.w == true) {
              removeSprite(cellCoords);
              cellCoords = {
                x: cellCoords.x - 1,
                y: cellCoords.y
              };
              drawSprite(cellCoords);
            }
            break;
          case 87:
          case 38: // north
            if (cell.n == true) {
              removeSprite(cellCoords);
              cellCoords = {
                x: cellCoords.x,
                y: cellCoords.y - 1
              };
              drawSprite(cellCoords);
            }
            break;
          case 68:
          case 39: // east
            if (cell.e == true) {
              removeSprite(cellCoords);
              cellCoords = {
                x: cellCoords.x + 1,
                y: cellCoords.y
              };
              drawSprite(cellCoords);
            }
            break;
          case 83:
          case 40: // south
            if (cell.s == true) {
              removeSprite(cellCoords);
              cellCoords = {
                x: cellCoords.x,
                y: cellCoords.y + 1
              };
              drawSprite(cellCoords);
            }
            break;
        }
      }
    
      this.bindKeyDown = function() {
        window.addEventListener("keydown", check, false);
    
        $("#view").swipe({
          swipe: function(
            event,
            direction,
            distance,
            duration,
            fingerCount,
            fingerData
          ) {
            console.log(direction);
            switch (direction) {
              case "up":
                check({
                  keyCode: 38
                });
                break;
              case "down":
                check({
                  keyCode: 40
                });
                break;
              case "left":
                check({
                  keyCode: 37
                });
                break;
              case "right":
                check({
                  keyCode: 39
                });
                break;
            }
          },
          threshold: 0
        });
      };
    
      this.unbindKeyDown = function() {
        window.removeEventListener("keydown", check, false);
        $("#view").swipe("destroy");
      };
    
      drawSprite(maze.startCoord());
    
      this.bindKeyDown();
    }
    
    var mazeCanvas = document.getElementById("mazeCanvas");
    var ctx = mazeCanvas.getContext("2d");
    var sprite;
    var finishSprite;
    var maze, draw, player;
    var cellSize;
    var difficulty;
    // sprite.src = 'media/sprite.png';
    
    window.onload = function() {
      let viewWidth = $("#view").width();
      let viewHeight = $("#view").height();
      if (viewHeight < viewWidth) {
        ctx.canvas.width = viewHeight - viewHeight / 100;
        ctx.canvas.height = viewHeight - viewHeight / 100;
      } else {
        ctx.canvas.width = viewWidth - viewWidth / 100;
        ctx.canvas.height = viewWidth - viewWidth / 100;
      }
    
      //Load and edit sprites
      var completeOne = false;
      var completeTwo = false;
      var isComplete = () => {
        if(completeOne === true && completeTwo === true)
           {
             console.log("Runs");
             setTimeout(function(){
               makeMaze();
             }, 500);         
           }
      };
      sprite = new Image();
      sprite.src =
        "https://image.ibb.co/dr1HZy/Pf_RWr3_X_Imgur.png" +
        "?" +
        new Date().getTime();
      sprite.setAttribute("crossOrigin", " ");
      sprite.onload = function() {
        sprite = changeBrightness(1.2, sprite);
        completeOne = true;
        console.log(completeOne);
        isComplete();
      };
    
      finishSprite = new Image();
      finishSprite.src = "https://image.ibb.co/b9wqnJ/i_Q7m_U25_Imgur.png"+
      "?" +
      new Date().getTime();
      finishSprite.setAttribute("crossOrigin", " ");
      finishSprite.onload = function() {
        finishSprite = changeBrightness(1.1, finishSprite);
        completeTwo = true;
        console.log(completeTwo);
        isComplete();
      };
      
    };
    
    window.onresize = function() {
      let viewWidth = $("#view").width();
      let viewHeight = $("#view").height();
      if (viewHeight < viewWidth) {
        ctx.canvas.width = viewHeight - viewHeight / 100;
        ctx.canvas.height = viewHeight - viewHeight / 100;
      } else {
        ctx.canvas.width = viewWidth - viewWidth / 100;
        ctx.canvas.height = viewWidth - viewWidth / 100;
      }
      cellSize = mazeCanvas.width / difficulty;
      if (player != null) {
        draw.redrawMaze(cellSize);
        player.redrawPlayer(cellSize);
      }
    };
    
    function makeMaze() {
      //document.getElementById("mazeCanvas").classList.add("border");
      if (player != undefined) {
        player.unbindKeyDown();
        player = null;
      }
      var e = document.getElementById("diffSelect");
      difficulty = e.options[e.selectedIndex].value;
      cellSize = mazeCanvas.width / difficulty;
      maze = new Maze(difficulty, difficulty);
      draw = new DrawMaze(maze, ctx, cellSize, finishSprite);
      player = new Player(maze, mazeCanvas, cellSize, displayVictoryMess, sprite);
      if (document.getElementById("mazeContainer").style.opacity < "100") {
        document.getElementById("mazeContainer").style.opacity = "100";
      }
    }
    
        </script>
      </body>
    </html>
    `))
	})

	r.GET("/tic_tac_toe", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`<!doctype html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport"
              content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Tic Tac Toe</title>
        <style>
        body {
          font-family: "Arial", sans-serif;
      }
      
      section {
          text-align: center;
      }
      
      .game--container {
          display: grid;
          grid-template-columns: repeat(3, auto);
          width: 306px;
          margin: 50px auto;
      }
      
      .cell {
          font-family: "Permanent Marker", cursive;
          width: 100px;
          height: 100px;
          box-shadow: 0 0 0 1px #333333;
          border: 1px solid #333333;
          cursor: pointer;
      
          line-height: 100px;
          font-size: 60px;
      }
        </style>
    </head>
    <body>
        <section>
            <h1 class="game--title">Tic Tac Toe</h1>
            <div class="game--container">
                <div data-cell-index="0" class="cell"></div>
                <div data-cell-index="1" class="cell"></div>
                <div data-cell-index="2" class="cell"></div>
                <div data-cell-index="3" class="cell"></div>
                <div data-cell-index="4" class="cell"></div>
                <div data-cell-index="5" class="cell"></div>
                <div data-cell-index="6" class="cell"></div>
                <div data-cell-index="7" class="cell"></div>
                <div data-cell-index="8" class="cell"></div>
            </div>
            <h2 class="game--status"></h2>
            <button class="game--restart">Restart Game</button>
        </section>
    
        <script>
        const statusDisplay = document.querySelector('.game--status');
let gameActive = true;
let currentPlayer = "X";
let gameState = ["", "", "", "", "", "", "", "", ""];
const winningMessage = () => `+"`"+`Player \${currentPlayer} has won!`+"`"+`;
const drawMessage = () => `+"`"+`Game ended in a draw!`+"`"+`;
const currentPlayerTurn = () => `+"`"+`It's \${currentPlayer}'s turn`+"`"+`;
statusDisplay.innerHTML = currentPlayerTurn();
const winningConditions = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6]
];
function handleCellPlayed(clickedCell, clickedCellIndex) {
    gameState[clickedCellIndex] = currentPlayer;
    clickedCell.innerHTML = currentPlayer;
}
function handlePlayerChange() {
    currentPlayer = currentPlayer === "X" ? "O" : "X";
    statusDisplay.innerHTML = currentPlayerTurn();
}
function handleResultValidation() {
    let roundWon = false;
    for (let i = 0; i <= 7; i++) {
        const winCondition = winningConditions[i];
        let a = gameState[winCondition[0]];
        let b = gameState[winCondition[1]];
        let c = gameState[winCondition[2]];
        if (a === '' || b === '' || c === '') {
            continue;
        }
        if (a === b && b === c) {
            roundWon = true;
            break
        }
    }
    if (roundWon) {
        statusDisplay.innerHTML = winningMessage();
        gameActive = false;
        return;
    }
    let roundDraw = !gameState.includes("");
    if (roundDraw) {
        statusDisplay.innerHTML = drawMessage();
        gameActive = false;
        return;
    }
    handlePlayerChange();
}
function handleCellClick(clickedCellEvent) {
    const clickedCell = clickedCellEvent.target;
    const clickedCellIndex = parseInt(clickedCell.getAttribute('data-cell-index'));
    if (gameState[clickedCellIndex] !== "" || !gameActive) {
        return;
    }
    handleCellPlayed(clickedCell, clickedCellIndex);
    handleResultValidation();
}
function handleRestartGame() {
    gameActive = true;
    currentPlayer = "X";
    gameState = ["", "", "", "", "", "", "", "", ""];
    statusDisplay.innerHTML = currentPlayerTurn();
    document.querySelectorAll('.cell').forEach(cell => cell.innerHTML = "");
}
document.querySelectorAll('.cell').forEach(cell => cell.addEventListener('click', handleCellClick));
document.querySelector('.game--restart').addEventListener('click', handleRestartGame);
        </script>
    </body>`))
	})

	r.GET("/clock", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
    
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Analogue Clock</title>
    <style>
    body{
        background-color: #000;
    }
    
    #canvas{
        margin-left: 300px;
        margin-top: 50px;
    }
    </style>
</head>
<body>
    <canvas id="canvas" width="600" height="600"></canvas>
    <script>
    //create a canvas object fro HTML element
var canvas = document.getElementById('canvas');
//create a 2d drawing object
var ctx = canvas.getContext('2d');
//calculate the clock radius by using the height
var radius = canvas.height / 2;
//remap the x and y axis to the center of the canvas
ctx.translate(radius, radius);
//reduce clock radius by 90%
radius = radius * 0.90;

setInterval(drawClock, 1000); //run the drawClock function every second.

function drawClock(){
    drawFace(ctx, radius);
    drawNumbers(ctx, radius);
    drawTime(ctx, radius);
}

function drawFace(ctx, radius){
    var grad;
    //draw white circle for the face
    ctx.beginPath();
    ctx.arc(0,0,radius,0,2*Math.PI);
    ctx.fillStyle = "White";
    ctx.fill();
    // create a radial gradient (inner, middle, and outer edge of clock)
    grad = ctx.createRadialGradient(0,0,radius*0.95, 0,0,radius*1.05);
    grad.addColorStop(0, '#333');
    grad.addColorStop(0.5, 'white');
    grad.addColorStop(1, '#333');
    //define gradient as stroke style
    ctx.strokeStyle = grad;
    ctx.lineWidth = radius*0.1;
    ctx.stroke(); 
    //draw the center of the clock
    ctx.beginPath();
    ctx.arc(0,0, radius*0.1,0,2*Math.PI);
    ctx.fillStyle = '#333';
    ctx.fill();
}

function drawNumbers(ctx, radius) {
    var ang;
    var num;
    ctx.font = radius*0.15 + "px arial"; //set font at 15\% \of radius
    ctx.textBaseline = "middle"; //set text alignment to middle
    ctx.textAlign = "center"; //set text alignment to center
    for(num=1; num < 13; num++){ //calculate the print position for each number
        ang = num *Math.PI /6;
        ctx.rotate(ang);
        ctx.translate(0, -radius*0.85);
        ctx.rotate(-ang);
        ctx.fillText(num.toString(), 0, 0);
        ctx.rotate(ang);
        ctx.translate(0, radius*0.85);
        ctx.rotate(-ang);
    }
}

function drawTime(ctx, radius){
    var now = new Date();
    var hour = now.getHours();
    var minute = now.getMinutes();
    var second = now.getSeconds();
    //hour
    hour = hour%12;
    //calculate angle of hour hand
    hour = (hour*Math.PI/6)+(minute*Math.PI/(6*60))+(second*Math.PI/(360*60));
    //make hour hand 50\% \of canvas's radius
    drawHand(ctx, hour, radius*0.5, radius*0.07);
    //minute
    //calculate angle of minute hand
    minute=(minute*Math.PI/30)+(second*Math.PI/(30*60));
    //make minute hand 80\% \of canvas's radius
    drawHand(ctx, minute, radius*0.8, radius*0.07);
    //second
    //calculate angle of second hand
    second=(second*Math.PI/30);
    //make second hand 90\% \of canvas's radius
    drawHand(ctx, second, radius*0.9, radius*0.02);
}

function drawHand(ctx, pos, length, width){
    ctx.beginPath();
    ctx.lineWidth = width;
    ctx.lineCap = "round";
    ctx.moveTo(0,0);
    ctx.rotate(pos);
    ctx.lineTo(0, -length);
    ctx.stroke();
    ctx.rotate(-pos);
}
    </script>
</body>
    `))
	})

	r.GET("/hangman", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
    <!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hangman</title>
    <style>
    * {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

html {
    font-size: 62.5%;
}

body {
    background: #2B292E;
    color: #fafafa;
    font-family: Helvetica, Arial, sans-serif;
    font-size: 1.6rem;
    display: flex;
    align-items: center;
    justify-content: center;
    min-height: 100vh;
}

span {
    border-bottom: 1px solid #534f59;
    display: inline-block;
    font-size: 2rem;
    height: 2.4rem;
    line-height: 2.4rem;
    margin: 0 .1rem;
    text-align: center;
    text-transform: uppercase;
    width: 2.4rem;
}

p {
    font-weight: 300;
    margin-bottom: .8rem;
}

.puzzle {
    display: flex;
    margin-bottom: 4.8rem;
}

.button {
    background: #7044a0;
    border: none;
    border-bottom: 2px solid #603a88;
    cursor: pointer;
    color: white;
    font-size: 1.4rem;
    font-weight: 300;
    padding: .8rem;
    transition: background .3s ease, color .3s ease;
}

.button:hover {
    background: #5F3A87;
}
    </style>
</head>
<body>
    <div>
        <div id="puzzle" class="puzzle"></div>
        <p id="guesses"></p>
        <button id="reset" class="button">Reset</button>
    </div>
    <script>
    const getPuzzle = async (wordCount) => {
        const response = await fetch(`+"`"+`https://puzzle.mead.io/puzzle?wordCount=\${wordCount}`+"`"+`)
            if (response.status === 200){
                const data = await response.json()
                return data.puzzle
            } else {
                throw new Error('Unable to fetch puzzle')
            }
    }
    class Hangman {
        constructor(word, remainingGuesses){
            this.word = word.toLowerCase().split('');
            this.remainingGuesses = remainingGuesses;
            this.guessedLetters = [];
            this.status = 'playing';
        }
    
        get puzzle() {
            let puzzle = '';
            this.word.forEach((letter) => {
            if (this.guessedLetters.includes(letter) || letter === ' '){
                puzzle += letter;
            } else {
                puzzle += '*'
            }
            })
            return puzzle;
        }
    
        makeGuess (guess){
            guess = guess.toLowerCase();
            const isUnique = !this.guessedLetters.includes(guess);
            const isBadGuess = !this.word.includes(guess);
            
        if (this.status !== 'playing'){
            return
        }
        
            if (isUnique){
                this.guessedLetters.push(guess)
            }
                
            if (isUnique && isBadGuess){
                this.remainingGuesses--
            }
            this.calculateStatus();
        }
    
        get statusMessage(){
            if (this.status === 'playing'){
                return `+"`"+`Guesses left: \${this.remainingGuesses}`+"`"+`
            } else if (this.status === 'failed') {
                return `+"`"+`Nice try! The word was "\${this.word.join('')}" `+"`"+`
            } else {
                return 'Great work! You guessed the word!'
            }
        }
    
        calculateStatus(){
            const finished = this.word.every((letter) => this.guessedLetters.includes(letter) || letter === ' ')
            
            if (this.remainingGuesses === 0){
                this.status = 'failed'
            } else if (finished){
                this.status = 'finished'
            } else {
                this.status = 'playing'
            }
        }
    
    }
    
    let game1
const puzzleDIV = document.querySelector('#puzzle');
const remainingDIV = document.querySelector('#guesses');

window.addEventListener('keypress', (e) => {

    const guess = String.fromCharCode(e.charCode);
    game1.makeGuess(guess);
    render()
})

const render = () => {
    puzzleDIV.innerHTML = ''
    remainingDIV.textContent = game1.statusMessage;

    game1.puzzle.split('').forEach((letter) => {
        const letterEl = document.createElement('span')
        letterEl.textContent = letter
        puzzleDIV.appendChild(letterEl)
    })
}

const startGame = async () => {
    const puzzle = await getPuzzle('3')
    game1 = new Hangman(puzzle, 5)
    render()
}

document.querySelector('#reset').addEventListener('click', startGame)
startGame()
    </script>
</body>
</html>
    `))
	})

	r.GET("/puzzles", func(c *gin.Context) {
		c.Data(http.StatusOK, ContentTypeHTML, []byte(`
<!DOCTYPE html>
<html>
<head>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jqueryui/1.12.1/jquery-ui.min.js"></script>
<title>Picture Puzzle</title>
<style>
body {
    font-family: 'Segoe UI', Calibri, Arial;
    margin:0;
    background-color: lightgreen;
}
h2{
    font-weight:normal;
    text-align:center;
}
h3{
    font-weight:normal;
    margin:3px 0px;
    text-align:center;
}
#collage hr{
border:none;
border-top:2px solid #f5f2f2;
height:1px;
}

#collage #playPanel {
background-color:#c2defc;
padding:10px 0px;
margin: 10px auto;
max-width:800px;
width:95%;
}

#collage #actualImageBox {
display: inline-block;
font-size:0.8em;
margin: 10px 10px;
vertical-align: top;
width:280px;
}

#collage #stepBox, #collage #timeBox {
display:inline-block;
width:48%;
}

#collage #stepBox div {
background-color:#c2defc;
display:inline-block;
padding:1px 4px;
margin: 0px auto;
max-width:800px;
}

#collage img#actualImage{
border:2px solid #a46;
height:280px;
width:280px;
}

#collage #sortable {
border:2px solid #a46;
list-style-type: none;
display: inline-block;
margin: 10px;
padding: 0;
width: 400px;
}

#collage #sortable li {
    background-size: 400% 400%;
    border: none;
    cursor: pointer;
    margin: 0;
    padding: 0;
    float: left;
    width: 100px;
    height: 100px;
}

#collage button  {
background-color:#f5f2f2;
border:1px solid #cce;
display: inline;
font-size: 14px;
height: auto;
width: auto;
padding: 3px 8px;
}
</style>

<script>
    var timerFunction;

var imagePuzzle = {
stepCount: 0,
startTime: new Date().getTime(),
startGame: function (images, gridSize) {
    this.setImage(images, gridSize);
    $('#playPanel').show();
    $('#sortable').randomize();
    this.enableSwapping('#sortable li');
    this.stepCount = 0;
    this.startTime = new Date().getTime();
    this.tick();
},
tick: function () {
    var now = new Date().getTime();
    var elapsedTime = parseInt((now - imagePuzzle.startTime) / 1000, 10);
    $('#timerPanel').text(elapsedTime);
    timerFunction = setTimeout(imagePuzzle.tick, 1000);
},
enableSwapping: function (elem) {
    $(elem).draggable({
        snap: '#droppable',
        snapMode: 'outer',
        revert: "invalid",
        helper: "clone"
    });
    $(elem).droppable({
        drop: function (event, ui) {
            var $dragElem = $(ui.draggable).clone().replaceAll(this);
            $(this).replaceAll(ui.draggable);

            currentList = $('#sortable > li').map(function (i, el) { return $(el).attr('data-value'); });
            if (isSorted(currentList))
                $('#actualImageBox').empty().html($('#gameOver').html());
            else {
                var now = new Date().getTime();
                imagePuzzle.stepCount++;
                $('.stepCount').text(imagePuzzle.stepCount);
                $('.timeCount').text(parseInt((now - imagePuzzle.startTime) / 1000, 10));
            }

            imagePuzzle.enableSwapping(this);
            imagePuzzle.enableSwapping($dragElem);
        }
    });
},

setImage: function (images, gridSize) {
    console.log(gridSize);
    gridSize = gridSize || 4; // If gridSize is null or not passed, default it as 4.
    console.log(gridSize);
    var percentage = 100 / (gridSize - 1);
    var image = images[Math.floor(Math.random() * images.length)];
    $('#imgTitle').html(image.title);
    $('#actualImage').attr('src', image.src);
    $('#sortable').empty();
    for (var i = 0; i < gridSize * gridSize; i++) {
        var xpos = (percentage * (i % gridSize)) + '%';
        var ypos = (percentage * Math.floor(i / gridSize)) + '%';
        var li = $('<li class="item" data-value="' + (i) + '"></li>').css({
            'background-image': 'url(' + image.src + ')',
            'background-size': (gridSize * 100) + '%',
            'background-position': xpos + ' ' + ypos,
            'width': 400 / gridSize,
            'height': 400 / gridSize
        });
        $('#sortable').append(li);
    }
    $('#sortable').randomize();
}
};

function isSorted(arr) {
for (var i = 0; i < arr.length - 1; i++) {
    if (arr[i] != i)
        return false;
}
return true;
}
$.fn.randomize = function (selector) {
var $elems = selector ? $(this).find(selector) : $(this).children(),
    $parents = $elems.parent();

$parents.each(function () {
    $(this).children(selector).sort(function () {
        return Math.round(Math.random()) - 0.5;
    }).remove().appendTo(this);
});
return this;
};
</script>
</head>
<body>
<div id="collage">
    <h2>Picture Puzzle</h2>
    <hr />
    <div id="playPanel" style="padding:5px;display:none;">
        <h3 id="imgTitle"></h3> <hr />
        <div style="display:inline-block; margin:auto; width:95%; vertical-align:top;">
            <ul id="sortable" class="sortable"></ul>
            <div id="actualImageBox">
                <div id="stepBox">
                    <div>Count:</div><div class="stepCount">0</div>
                </div>
                <div id="timeBox">
                    Time: <span id="timerPanel"></span> seconds
                </div>
                <img id="actualImage" />
                <div>See the picture and solve the puzzle.</div>
                <p id="levelPanel">
                    <input type="radio" name="level" id="easy" checked="checked" value="3" /> <label for="easy">Easy</label>
                    <input type="radio" name="level" id="medium" value="4" /> <label for="medium">Medium</label>
                    <input type="radio" name="level" id="hard" value="5" /> <label for="hard">Hard</label>
                    <input type="radio" name="level" id="veryhard" value="6" /> <label for="veryhard">Very Hard</label>
                </p>
                <div>
                    <button id="btnRule" type="button" class="btn" onclick="rules();">Rules</button>
                    <button id="newPhoto" type="button" class="btn">Next Picture</button>
                </div>
            </div>
        </div>
    </div>
    <div id="gameOver" style="display:none;">
        <div style="background-color: #fc9e9e; padding: 5px 10px 20px 10px; text-align: center; ">
            <h2 style="text-align:center">Game over !!</h2>
            Congratulations!! <br /> You got the picture ready.
            <br />
            Time: <span class="stepCount">0</span> gear.
            <br />
            Time spent: <span class="timeCount">0</span> seconds<br />
            <div>
                <button type="button" onclick="window.location.reload(true);">Play more</button>
            </div>
        </div>
    </div>

    <script>
    var images = [
        { src: 'https://raw.githubusercontent.com/sreegithub19/JavaScript-Applications/main/public/puzzle/img/neeruti.jpg', title: 'Neeruti manor' },
        { src: 'https://raw.githubusercontent.com/sreegithub19/JavaScript-Applications/main/public/puzzle/img/harju_madise.jpg', title: 'Harju-Madis Church' },
        { src: 'https://raw.githubusercontent.com/sreegithub19/JavaScript-Applications/main/public/puzzle/img/rahumae.jpg', title: 'Rahumäe train station' },
        { src: 'https://raw.githubusercontent.com/sreegithub19/JavaScript-Applications/main/public/puzzle/img/kakumae.jpg', title: 'Kakumäe Harbor' },
        { src: 'https://raw.githubusercontent.com/sreegithub19/JavaScript-Applications/main/public/puzzle/img/kohila.jpg', title: 'Kohila mill' }
    ];
    
    $(function () {
        var gridSize = $('#levelPanel :radio:checked').val();
        imagePuzzle.startGame(images, gridSize);
        $('#newPhoto').click(function () {
            var gridSize = $('#levelPanel :radio:checked').val();
            imagePuzzle.startGame(images, gridSize);
        });
    
        $('#levelPanel :radio').change(function (e) {
            var gridSize = $(this).val();
            imagePuzzle.startGame(images, gridSize);
        });
    });
    function rules() {
        alert('Rearrange the pieces so that you get a sample image. The steps taken are counted');
    }
    </script>
    </div>
</body>
</html>
`))
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)

}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
