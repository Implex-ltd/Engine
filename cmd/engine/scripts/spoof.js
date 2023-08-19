delete Object.getPrototypeOf(navigator).webdriver;
navigator.webdriver = false;

const originalHardwareConcurrency = navigator.hardwareConcurrency;
const originalNavigator = Object.create(navigator);

const availableExtensions = ["EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_sRGB", "EXT_texture_filter_anisotropic", "OES_element_index_uint", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_astc", "WEBGL_compressed_texture_atc", "WEBGL_compressed_texture_etc", "WEBGL_compressed_texture_etc1", "WEBGL_compressed_texture_pvrtc", "WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBGL_multiview", "WEBGL_texture_compression_bptc", "WEBGL_texture_compression_rgtc", "WEBGL_texture_compression_s3tc", "WEBGL_texture_compression_s3tc_srgb", "WEBGL_texture_filter_anisotropic"];

function addNoiseToCanvas() {
    const originalFillText = CanvasRenderingContext2D.prototype.fillText;
    CanvasRenderingContext2D.prototype.fillText = function (text, x, y, maxWidth) {
        x += Math.random() * 2 - 1;
        y += Math.random() * 2 - 1;
        originalFillText.call(this, text, x, y, maxWidth);
    };
}

function webgl(vendor, renderer) {
    // Store a reference to the original getParameter function
    const originalGetParameter = WebGLRenderingContext.prototype.getParameter;
    const originalGetParameter2 = WebGL2RenderingContext.prototype.getParameter;

    WebGL2RenderingContext.prototype.getParameter = function (parameter) {
        if (parameter === this.RENDERER) {
            return vendor;
        } else if (parameter === this.VENDOR) {
            return renderer;
        } else if (parameter === this.EXTENSIONS) {
            // Modify this part according to your requirements
            const numRandomExtensions = Math.floor(Math.random() * (availableExtensions.length + 1));
            const randomExtensions = [];
            while (randomExtensions.length < numRandomExtensions) {
                const randomIndex = Math.floor(Math.random() * availableExtensions.length);
                const randomExtension = availableExtensions[randomIndex];
                if (!randomExtensions.includes(randomExtension)) {
                    randomExtensions.push(randomExtension);
                }
            }
            return randomExtensions.join(" ");
        }

        // If the parameter is not overridden, use the original function
        return originalGetParameter2.call(this, parameter);
    };

    // Override the getParameter function
    WebGLRenderingContext.prototype.getParameter = function (parameter) {
        if (parameter === this.RENDERER) {
            return vendor;
        } else if (parameter === this.VENDOR) {
            return renderer;
        } else if (parameter === this.EXTENSIONS) {
            // Modify this part according to your requirements
            const numRandomExtensions = Math.floor(Math.random() * (availableExtensions.length + 1));
            const randomExtensions = [];
            while (randomExtensions.length < numRandomExtensions) {
                const randomIndex = Math.floor(Math.random() * availableExtensions.length);
                const randomExtension = availableExtensions[randomIndex];
                if (!randomExtensions.includes(randomExtension)) {
                    randomExtensions.push(randomExtension);
                }
            }
            return randomExtensions.join(" ");
        }

        // If the parameter is not overridden, use the original function
        return originalGetParameter.call(this, parameter);
    };
}

function devices() {
    const originalEnumerateDevices = navigator.mediaDevices.enumerateDevices;

    navigator.mediaDevices.enumerateDevices = async function () {
        const originalDevices = await originalEnumerateDevices.call(navigator.mediaDevices);
        const fakeDevices = [{
            kind: "audioinput",
            label: "Fake Microphone",
            deviceId: "fake-mic-id"
        }, {
            kind: "videoinput",
            label: "Fake Webcam",
            deviceId: "fake-webcam-id"
        }, {
            kind: "audioinput",
            label: "USB Microphone",
            deviceId: "usb-mic-id"
        }, {
            kind: "audioinput",
            label: "External Microphone",
            deviceId: "external-mic-id"
        }, {
            kind: "videoinput",
            label: "Built-in Camera",
            deviceId: "built-in-camera-id"
        }, {
            kind: "videoinput",
            label: "External Webcam",
            deviceId: "external-webcam-id"
        }];
        return originalDevices.concat(fakeDevices);
    };
}

function fonts() {
    const originalFontGetter = Object.getOwnPropertyDescriptor(HTMLElement.prototype, 'style').get;

    Object.defineProperty(HTMLElement.prototype, 'style', {
        get: function () {
            const fontList = [
                // List of spoofed fonts
                'Arial', 'Comic Sans MS', 'Times New Roman', 'Courier New', 'Verdana'
                // Add your desired fonts here
            ];

            const randomFontIndex = Math.floor(Math.random() * fontList.length);
            const randomFont = fontList[randomFontIndex];

            const originalStyle = originalFontGetter.call(this);
            originalStyle.fontFamily = `"${randomFont}", ${originalStyle.fontFamily}`;

            return originalStyle;
        },
    });
}

function webrtc() {
    window.RTCPeerConnection = function () {
        throw new Error('WebRTC is disabled');
    };
}

// Define the fake plugins data
const fakePlugins = [{
    name: "Chrome PDF Viewer",
    filename: "internal-pdf-viewer",
    description: ""
}, {
    name: "Shockwave Flash",
    filename: "pepflashplayer",
    description: "Shockwave Flash 99.0 r0"
}, {
    name: "Java(TM) Platform",
    filename: "java-runtime-environment",
    description: "Next Generation Java Plug-in for Mozilla browsers"
}, {
    name: "Widevine Content Decryption Module",
    filename: "widevinecdm",
    description: "Enables Widevine licenses for playback of HTML audio/video content"
}];

function getRandomNumber(min, max) {
    const range = Math.floor((max - min + 1) / 2);
    const randomOffset = Math.floor(Math.random() * range);
    return randomOffset * 2 + min;
}

/*
let oof = {
    languages: {
        get: function () {
            return ["en-US", "en"];
        }
    },
    language: {
        get: function () {
            return "en-US";
        }
    },
    hardwareConcurrency: {
        get: function () {
            return getRandomNumber(2, 8);
        }
    },
    deviceMemory: {
        get: function () {
            return getRandomNumber(4, 8);
        },
        configurable: false,
        enumerable: true
    },
    plugins: {
        get: function () {
            return fakePlugins.map(plugin => ({
                ...plugin,
                length: fakePlugins.length
            }));
        }
    }
}

// Override navigator properties
Object.defineProperties(navigator, oof);
Object.defineProperties(Navigator.prototype, oof);
Object.defineProperties(Navigator, oof);
*/

function spoofall() {
    addNoiseToCanvas();
    webgl();
    //fonts();
}

// one time
webrtc();
devices();

spoofall();