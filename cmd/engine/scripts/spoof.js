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

Object.defineProperty(performance, "now", {
    value: function () {
        return 14112.800000000745 + (Math.random() * 7000.000000000000);
    }
});

Object.defineProperty(SVGRect.prototype, "x", { value: Math.floor(Math.random() * 100) + 111 })
Object.defineProperty(SVGRect.prototype, "y", { value: Math.floor(Math.random() * 100) + 111 })

Object.defineProperty(SVGRect.prototype, "height", { value: Math.floor(Math.random() * 100) + 111 })
Object.defineProperty(SVGRect.prototype, "width", { value: Math.floor(Math.random() * 100) + 111 })

Object.defineProperty(SVGRectElement.prototype, "getBBox", { value: Math.floor(Math.random() * 100) + 111 })

Object.defineProperty(SVGTextContentElement.prototype, "getSubStringLength", { value: () => Math.floor(Math.random() * 100) + 111 })
Object.defineProperty(SVGTextContentElement.prototype, "getComputedTextLength", { value: () => Math.floor(Math.random() * 100) + 111 })

function spoofall() {
    addNoiseToCanvas();
}
spoofall();