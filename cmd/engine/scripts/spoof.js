delete Object.getPrototypeOf(navigator).webdriver;
if (navigator && navigator.webdriver !== undefined) {
    Object.defineProperty(navigator, "webdriver", { value: false });
}

const randomOffset = (x) => Math.random() * x - 1;

const originalFillText = CanvasRenderingContext2D.prototype.fillText;
const originalDrawImage = CanvasRenderingContext2D.prototype.drawImage;
const originalfillRect = CanvasRenderingContext2D.prototype.fillRect;

CanvasRenderingContext2D.prototype.fillText = function (text, x, y, maxWidth) {
    x += randomOffset(2);
    y += randomOffset(2);
    originalFillText.call(this, text, x, y, maxWidth);
};

CanvasRenderingContext2D.prototype.drawImage = function (img, sx, sy, sw, sh, dx, dy, dw, dh) {
    dx += randomOffset(4);
    dy += randomOffset(4);
    originalDrawImage.call(this, img, sx, sy, sw, sh, dx, dy, dw, dh);
}

CanvasRenderingContext2D.prototype.fillRect = function (x, y, w, h) {
    x += randomOffset(2);
    y += randomOffset(2);
    originalfillRect.call(this, x, y, w, h);
};