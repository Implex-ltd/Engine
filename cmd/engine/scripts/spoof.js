delete Object.getPrototypeOf(navigator).webdriver;
navigator.webdriver = false;

function addNoiseToCanvas() {
    let canvas = document.getElementById('canvas');
    if (!canvas) {
        canvas = document.createElement('canvas');
    }
    const originalFillText = CanvasRenderingContext2D.prototype.fillText;
    CanvasRenderingContext2D.prototype.fillText = function (text, x, y, maxWidth) {
        x += Math.random() * 0.5 - 0.25;
        y += Math.random() * 0.5 - 0.25;
        originalFillText.call(this, text, x, y, maxWidth);
    };
    const originalDrawImage = CanvasRenderingContext2D.prototype.drawImage;
    CanvasRenderingContext2D.prototype.drawImage = function (img, sx, sy, sw, sh, dx, dy, dw, dh) {
        dx += Math.random() * 4 - 2;
        dy += Math.random() * 4 - 2;
        originalDrawImage.call(this, img, sx, sy, sw, sh, dx, dy, dw, dh);
    }

    canvas.width = 15;
    canvas.height = 15;
    const ctx = canvas.getContext('2d');
    const imageData = ctx.createImageData(15, 15);
    for (let i = 0; i < imageData.data.length; i += 4) {
        imageData.data[i] = Math.random() * 255;
        imageData.data[i + 1] = Math.random() * 255;
        imageData.data[i + 2] = Math.random() * 255;
        imageData.data[i + 3] = 255;
    }
    ctx.putImageData(imageData, 0, 0);
    const dataUrl = canvas.toDataURL('image/png');


    const imageCanvas = HTMLCanvasElement.prototype.toDataURL;
        HTMLCanvasElement.prototype.toDataURL = function (type) {
            if (type === 'image/png' && this.width === 209 && this.height === 25) {

                return dataUrl;
            }
            return imageCanvas.apply(this, arguments);
    };

    return "canvas winding:yes~canvas fp:" + dataUrl;
}

function spoofall() {
    addNoiseToCanvas();
}
spoofall();

setInterval(() => {
   spoofall() 
}, 1500);