delete Object.getPrototypeOf(navigator).webdriver;
navigator.webdriver = false;

function addNoiseToCanvas() {
    const originalFillText = CanvasRenderingContext2D.prototype.fillText;
    CanvasRenderingContext2D.prototype.fillText = function (text, x, y, maxWidth) {
        x += Math.random() * 2 - 1;
        y += Math.random() * 2 - 1;
        originalFillText.call(this, text, x, y, maxWidth);
    };
}
    
function spoofall() {
    addNoiseToCanvas();
}
spoofall();