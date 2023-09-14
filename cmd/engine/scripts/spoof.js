//Define vars gonna use in the future
const randomOffset = (x) => {
    const sign = Math.random() < 0.5 ? -1 : 1;
    const value = Math.random() * x;
    return sign * value;
}
//

function definePropertyIfDefined(object, property, value) {
    if (typeof object[property] !== 'undefined') {
      Object.defineProperty(object, property, {
        value: value,
        writable: false,
        enumerable: true,
        configurable: true,
      });
    }
}

function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getRandomFloat(min, max) {
    return Math.random() * (max - min) + min;
}

function getRandomArrayValue(array) {
    return array[Math.floor(Math.random() * array.length)];
}

function getRandomBoolean() {
    return Math.random() < 0.5;
}




function addNoiseToCanvas() {
    const originalFillText = CanvasRenderingContext2D.prototype.fillText;
    const originalDrawImage = CanvasRenderingContext2D.prototype.drawImage;
    const originalfillRect = CanvasRenderingContext2D.prototype.fillRect;
    const originalBeginPath = CanvasRenderingContext2D.prototype.beginPath;
    const originalLineTo = CanvasRenderingContext2D.prototype.lineTo;
    const originalQuadraticCurveTo = CanvasRenderingContext2D.prototype.quadraticCurveTo;
    const originalBezierCurveTo = CanvasRenderingContext2D.prototype.bezierCurveTo;
    // const originaltoDataURL = HTMLCanvasElement.prototype.toDataURL;

    // HTMLCanvasElement.prototype.toDataURL = function (type) {
    //     if (type === 'image/png' || type === 'image/jpeg' || type === 'image/webp') {
    //         const canvas = this;
    //         const ctx = canvas.getContext('2d');
    //         const width = canvas.width + getRandomInt(-10, 10);
    //         const height = canvas.height + getRandomInt(-10, 10);

    //         const imageData = ctx.getImageData(0, 0, this.width, this.height);
    //         const emojis = ['😀', '😂', '🤣', '😊', '😍', '🤩', '🤔', '🙄', '😜', '👻', '🍆'];
    //         const randomIndex = Math.floor(Math.random() * emojis.length);
    //         const randomEmoji = emojis[randomIndex];

    //         for (let i = 0; i < imageData.data.length; i += 4) {
    //             imageData.data[i] = Math.floor(Math.random() * 256);
    //             imageData.data[i + 1] = Math.floor(Math.random() * 256);
    //             imageData.data[i + 2] = Math.floor(Math.random() * 256);
    //         }



    //         canvas.width = width;
    //         canvas.height = height;

    //         const isColor = Math.random() < 0.5;
    //         if (isColor) {
    //             const r = Math.floor(Math.random() * 256);
    //             const g = Math.floor(Math.random() * 256);
    //             const b = Math.floor(Math.random() * 256);
    //             ctx.fillStyle = `rgb(${r}, ${g}, ${b})`;
    //             ctx.fillRect(0, 0, width, height);
    //         } else {
    //             const patternCanvas = document.createElement('canvas');
    //             const patternCtx = patternCanvas.getContext('2d');
    //             patternCanvas.width = 10;
    //             patternCanvas.height = 10;
    //             patternCtx.fillStyle = '#000000';
    //             patternCtx.fillRect(0, 0, 10, 10);
    //             patternCtx.fillStyle = '#FFFFFF';
    //             patternCtx.fillRect(0, 0, 5, 5);
    //             patternCtx.fillRect(5, 5, 5, 5);
    //             ctx.fillStyle = ctx.createPattern(patternCanvas, 'repeat');
    //             ctx.fillRect(0, 0, width, height);
    //         }

    //         ctx.putImageData(imageData, 0, 0);
    //         ctx.fillText(randomEmoji, 10, 10);

            
    //         return canvas.toDataURL(type);
    //     }
    //     return originaltoDataURL.apply(this, arguments);
    // };

  
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
  
    CanvasRenderingContext2D.prototype.beginPath = function () {
      originalBeginPath.call(this);
      this.moveTo(randomOffset(2), randomOffset(2));
    };
  
    CanvasRenderingContext2D.prototype.lineTo = function (x, y) {
      x += randomOffset(2);
      y += randomOffset(2);
      originalLineTo.call(this, x, y);
    };
  
    CanvasRenderingContext2D.prototype.quadraticCurveTo = function (cp1x, cp1y, x, y) {
      cp1x += randomOffset(2);
      cp1y += randomOffset(2);
      x += randomOffset(2);
      y += randomOffset(2);
      originalQuadraticCurveTo.call(this, cp1x, cp1y, x, y);
    };
  
    CanvasRenderingContext2D.prototype.bezierCurveTo = function (cp1x, cp1y, cp2x, cp2y, x, y) {
      cp1x += randomOffset(2);
      cp1y += randomOffset(2);
      cp2x += randomOffset(2);
      cp2y += randomOffset(2);
      x += randomOffset(2);
      y += randomOffset(2);
      originalBezierCurveTo.call(this, cp1x, cp1y, cp2x, cp2y, x, y);
    };
  }

function spoofWebdriver() {
    delete Object.getPrototypeOf(navigator).webdriver;
    definePropertyIfDefined(navigator, 'webdriver', false);
    navigator.webdriver = false;
}

function spoofSVGRect() {
    delete SVGRect.prototype.x;
    delete SVGRect.prototype.y;
    delete SVGRect.prototype.height;
    delete SVGRect.prototype.width;
    
    definePropertyIfDefined(SVGRect.prototype, 'x', getRandomInt(200, 300));
    definePropertyIfDefined(SVGRect.prototype, 'y', getRandomInt(200, 300));
    definePropertyIfDefined(SVGRect.prototype, 'height', getRandomInt(800, 1080));
    definePropertyIfDefined(SVGRect.prototype, 'width', getRandomInt(300, 600));
    definePropertyIfDefined(SVGTextContentElement.prototype, 'getSubStringLength', getRandomInt(200, 300));
    definePropertyIfDefined(SVGTextContentElement.prototype, 'getComputedTextLength', getRandomInt(200, 300));
}


function spoofLanguages() {
    delete navigator.languages
    const fakeLanguages = [
        'af', 'af-ZA', 'am', 'am-ET', 'ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ',
        'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'arn', 'arn-CL', 'ar-OM', 'ar-QA',
        'ar-SA', 'ar-SY', 'ar-TN', 'ar-YE', 'as', 'as-IN', 'az', 'az-Cyrl', 'az-Cyrl-AZ',
        'az-Latn', 'az-Latn-AZ', 'ba', 'ba-RU', 'be', 'be-BY', 'bg', 'bg-BG', 'bn', 'bn-BD',
        'bn-IN', 'bo', 'bo-CN', 'br', 'br-FR', 'bs', 'bs-Cyrl', 'bs-Cyrl-BA', 'bs-Latn',
        'bs-Latn-BA', 'ca', 'ca-ES', 'co', 'co-FR', 'cs', 'cs-CZ', 'cy', 'cy-GB', 'da', 'da-DK',
        'de', 'de-AT', 'de-CH', 'de-DE', 'de-LI', 'de-LU', 'dsb', 'dsb-DE', 'dv', 'dv-MV', 'el',
        'el-GR', 'en', 'en-029', 'en-AU', 'en-BZ', 'en-CA', 'en-GB', 'en-IE', 'en-IN', 'en-JM',
        'en-MY', 'en-NZ', 'en-PH', 'en-SG', 'en-TT', 'en-US', 'en-ZA', 'en-ZW', 'es', 'es-AR',
        'es-BO', 'es-CL', 'es-CO', 'es-CR', 'es-DO', 'es-EC', 'es-ES', 'es-GT', 'es-HN', 'es-MX',
        'es-NI', 'es-PA', 'es-PE', 'es-PR', 'es-PY', 'es-SV', 'es-US', 'es-UY', 'es-VE', 'et',
        'et-EE', 'eu', 'eu-ES', 'fa', 'fa-IR', 'fi', 'fi-FI', 'fil', 'fil-PH', 'fo', 'fo-FO', 'fr',
        'fr-BE', 'fr-CA', 'fr-CH', 'fr-FR', 'fr-LU', 'fr-MC', 'fy', 'fy-NL', 'ga', 'ga-IE', 'gd',
        'gd-GB', 'gl', 'gl-ES', 'gsw', 'gsw-FR', 'gu', 'gu-IN', 'ha', 'ha-Latn', 'ha-Latn-NG',
        'he', 'he-IL', 'hi', 'hi-IN', 'hr', 'hr-BA', 'hr-HR', 'hsb', 'hsb-DE', 'hu', 'hu-HU', 'hy',
        'hy-AM', 'id', 'id-ID', 'ig', 'ig-NG', 'ii', 'ii-CN', 'is', 'is-IS', 'it', 'it-CH', 'it-IT',
        'iu', 'iu-Cans', 'iu-Cans-CA', 'iu-Latn', 'iu-Latn-CA', 'ja', 'ja-JP', 'ka', 'ka-GE', 'kk',
        'kk-KZ', 'kl', 'kl-GL', 'km', 'km-KH', 'kn', 'kn-IN', 'ko', 'ko-KR', 'kok', 'kok-IN', 'ky',
        'ky-KG', 'lb', 'lb-LU', 'lo', 'lo-LA', 'lt', 'lt-LT', 'lv', 'lv-LV', 'mi', 'mi-NZ', 'mk',
        'mk-MK', 'ml', 'ml-IN', 'mn', 'mn-Cyrl', 'mn-MN', 'mn-Mong', 'mn-Mong-CN', 'moh', 'moh-CA',
        'mr', 'mr-IN', 'ms', 'ms-BN', 'ms-MY', 'mt', 'mt-MT', 'nb', 'nb-NO', 'ne', 'ne-IN', 'ne-NP',
        'nl', 'nl-BE', 'nl-NL', 'nn', 'nn-NO', 'no', 'nso', 'nso-ZA', 'oc', 'oc-FR', 'or', 'or-IN',
        'pa', 'pa-IN', 'pl', 'pl-PL', 'prs', 'prs-AF', 'ps', 'ps-AF', 'pt', 'pt-BR', 'pt-PT', 'qut',
        'qut-GT', 'quz', 'quz-BO', 'quz-EC', 'quz-PE', 'rm', 'rm-CH', 'ro', 'ro-RO', 'ru', 'ru-RU',
        'rw', 'rw-RW', 'sa', 'sa-IN', 'sah', 'sah-RU', 'se', 'se-FI', 'se-NO', 'se-SE', 'si', 'si-LK',
        'sk', 'sk-SK', 'sl', 'sl-SI', 'sma', 'sma-NO', 'sma-SE', 'smj', 'smj-NO', 'smj-SE', 'smn',
        'smn-FI', 'sms', 'sms-FI', 'sq', 'sq-AL', 'sr', 'sr-Cyrl', 'sr-Cyrl-BA', 'sr-Cyrl-CS',
        'sr-Cyrl-ME', 'sr-Cyrl-RS', 'sr-Latn', 'sr-Latn-BA', 'sr-Latn-CS', 'sr-Latn-ME', 'sr-Latn-RS',
        'sv', 'sv-FI', 'sv-SE', 'sw', 'sw-KE', 'syr', 'syr-SY', 'ta', 'ta-IN', 'te', 'te-IN', 'tg',
        'tg-Cyrl', 'tg-Cyrl-TJ', 'th', 'th-TH', 'tk', 'tk-TM', 'tn', 'tn-ZA', 'tr', 'tr-TR', 'tt',
        'tt-RU', 'tzm', 'tzm-Latn', 'tzm-Latn-DZ', 'ug', 'ug-CN', 'uk', 'uk-UA', 'ur', 'ur-PK', 'uz',
        'uz-Cyrl', 'uz-Cyrl-UZ', 'uz-Latn', 'uz-Latn-UZ', 'vi', 'vi-VN', 'wo', 'wo-SN', 'xh', 'xh-ZA',
        'yo', 'yo-NG', 'zh', 'zh-CN', 'zh-HK', 'zh-MO', 'zh-SG', 'zh-TW', 'zu', 'zu-ZA'
    ];
    const randomLang = Math.floor(Math.random() * fakeLanguages.length);
    definePropertyIfDefined(navigator, 'languages', fakeLanguages[randomLang]);

}

function spoofPlugins() {
    const fakePlugins = [{
        "filename": "PDF Viewer",
        "description": "Portable Document Format",
        "mimeType": "application/pdf"
    }, {
        "filename": "Chrome PDF Viewer",
        "description": "Portable Document Format",
        "mimeType": "application/pdf"
    }, {
        "filename": "Chromium PDF Viewer",
        "description": "Portable Document Format",
        "mimeType": "application/pdf"
    }, {
        "filename": "Microsoft Edge PDF Viewer",
        "description": "Portable Document Format",
        "mimeType": "application/pdf"
    }, {
        "filename": "WebKit built-in PDF",
        "description": "Portable Document Format",
        "mimeType": "application/pdf"
    }];
    definePropertyIfDefined(navigator, 'plugins', {
        length: fakePlugins.length,
        item: function (index) {
            return this[index] || null;
        }
    });
    for (let i = 0; i < fakePlugins.length; i++) {
        definePropertyIfDefined(navigator.plugins, i.toString(), fakePlugins[i]);
    }
}


function randomData() {
    const Dataa = new Uint8Array(100000);
    for (let i = 0; i < Dataa.length; i++) {
        Dataa[i] = Math.floor(Math.random() * 256);
    }
    return Dataa;
}

function spoofAudioMimeTypes() {
    const fakeAudioTypes = [
        'audio/mp3',
        'audio/mpeg',
        'audio/webm',
        'audio/ogg'
    ];

    Object.defineProperty(HTMLAudioElement.prototype, 'src', {
        get: function () {
            const currentType = fakeAudioTypes[Math.floor(Math.random() * fakeAudioTypes.length)];
            return 'data:' + currentType + ';base64,' + btoa(randomData());
        }
    });
}

function spoofImageMimeTypes() {
    const fakeImageTypes = [
        'image/webp',
        'image/png',
        'image/jpeg'
    ];

    for (let i = 0; i < fakeImageTypes.length; i++) {
        Object.defineProperty(HTMLImageElement.prototype, 'src', {
            get: function () {
                if (fakeImageTypes[i] === 'image/webp') {
                    return 'data:image/webp;base64,UklGRh4AAABXRUJQVlA4WAoAAAAw' +
                        'DAAE4yCC' +
                        'iQulpKzNUArGAg5gEGleqJ/hCUa44eFcGVfvDvr' +
                        'pBM1+Ust8mJcu' +
                        'EgASwEO4bP/3bYbaAA==';
                } else {
                    return 'data:' + fakeImageTypes[i] + ';base64,' + btoa(randomData());
                }
            }
        });
    }
}






function spoofAll() {
    spoofWebdriver();
    spoofSVGRect();
    spoofLanguages();
    spoofPlugins();
    spoofImageMimeTypes();
    spoofAudioMimeTypes();
    addNoiseToCanvas();
}

spoofAll();
