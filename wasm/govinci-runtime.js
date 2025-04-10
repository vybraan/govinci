// govinci-runtime.js

const Govinci = (() => {
    let rootElement = null;
    const callbackMap = {};
    const DEBUG = true;

    function renderNode(node, path = "") {
        const el = createElement(node);
        el.setAttribute("data-node-path", path);
        if (DEBUG) console.log("adding render node path", path);

        if (node.Type === "Spacer" && node.Props && node.Props.size) {
            el.style.height = `${node.Props.size}px`;
        }

        if (node.Children) {
            node.Children.forEach((child, i) => {
                const childEl = renderNode(child, `${path}/${i}`);
                el.appendChild(childEl);
            });
        }

        return el;
    }

    function createElement(node) {
        const el = document.createElement(tagForType(node.Type));

        if (node.Style) {
            Object.assign(el.style, styleFromGovinci(node.Style));
        }

        if (node.Props) {
            for (const [key, value] of Object.entries(node.Props)) {
                if (key.startsWith("on")) {
                    const event = mapEventName(key);
                    const existing = el.dataset[`listener_${key}`];
                    if (existing && callbackMap[existing]) {
                        el.removeEventListener(event, callbackMap[existing]);
                    }
                    const handler = (e) => {
                        const payload = extractEventPayload(e, node.Type);
                        if (DEBUG) console.log("Triggering:", value, payload);
                        window.GoInvokeCallback(value, payload);
                    };
                    el.addEventListener(event, handler);
                    el.dataset[`listener_${key}`] = value;
                    callbackMap[value] = handler;
                } else if (key === "value") {
                    el.value = value;
                } else if (key === "placeholder") {
                    el.placeholder = value;
                } else if (key === "content") {
                    el.textContent = value;
                }
                else if (key === "label") {
                    el.textContent = value;
                }

            }
        }

        return el;
    }

    function styleFromGovinci(style) {
        const out = {};
        if (style.FontSize) out.fontSize = `${style.FontSize}px`;
        if (style.TextColor) out.color = style.TextColor;
        if (style.Background) out.background = style.Background;
        if (style.Padding) out.padding = edgeToCSS(style.Padding);
        if (style.Margin) out.margin = edgeToCSS(style.Margin);
        if (style.BorderRadius) out.borderRadius = `${style.BorderRadius}px`;
        return out;
    }

    function edgeToCSS(edge) {
        return `${edge.Top}px ${edge.Right}px ${edge.Bottom}px ${edge.Left}px`;
    }

    function tagForType(type) {
        switch (type) {
            case "Text": return "span";
            case "Input":
            case "InputPassword":
            case "NumericInput":
            case "Checkbox": return "input";
            case "TextArea": return "textarea";
            case "Button": return "button";
            case "Image": return "img";
            case "Card":
            case "Row":
            case "Column":
            case "Scroll":
            case "SafeArea":
            case "Fragment":
            case "Spacer": return "div";
            default: return "div";
        }
    }

    function mapEventName(propKey) {
        return {
            onClick: "click",
            onChange: "input",
            onToggle: "change"
        }[propKey] || propKey.toLowerCase().replace(/^on/, "");
    }

    function extractEventPayload(e, type) {
        if (["input", "textarea", "numericinput", "inputpassword"].includes(type)) {
            return { value: e.target.value };
        }
        if (type === "checkbox") {
            return { value: e.target.checked };
        }
        return {};
    }


    function mount(jsonTree, mountPointId = "app") {
        const tree = typeof jsonTree === "string" ? JSON.parse(jsonTree) : jsonTree;
        const root = renderNode(tree, "root");
        rootElement = document.getElementById(mountPointId);
        rootElement.innerHTML = "";
        rootElement.appendChild(root);
    }

    function patch(patchList) {
        const patches = typeof patchList === "string" ? JSON.parse(patchList) : patchList;

        patches.forEach(p => {
            if (DEBUG) console.log("found patch in runtime patcher", p);
            const el = document.querySelector(`[data-node-path="${p.TargetID}"]`);
            if (!el) {
                if (DEBUG) console.log("could not find el to patch");
                return;
            }
            console.log("the element is:",el)

            switch (p.Type) {
                case "update-props":
                    for (const [k, v] of Object.entries(p.Changes)) {
                        if (k === "value") el.value = v;
                        else if (k === "content") el.textContent = v;
                        else if (k === "placeholder") el.placeholder = v;
                        else if (k.startsWith("on")) {
                            const event = mapEventName(k);
                            const oldListenerId = el.dataset[`listener_${k}`];

                            // 🔥 Remove o antigo handler
                            if (oldListenerId && callbackMap[oldListenerId]) {
                                el.removeEventListener(event, callbackMap[oldListenerId]);
                                delete callbackMap[oldListenerId];
                            }
                            // 🧠 Registra novo handler com novo ID
                            const handler = (e) => {
                                const payload = extractEventPayload(e, el.tagName.toLowerCase());
                                if (DEBUG) console.log("Triggering:", v, payload);
                                window.GoInvokeCallback(v, payload);
                            };


                            el.addEventListener(event, handler);
                            el.dataset[`listener_${k}`] = v;
                            callbackMap[v] = handler;
                        }
                    }
                    break;


                case "update-style":
                    Object.assign(el.style, styleFromGovinci(p.Changes));
                    break;

                case "replace":
                    const newEl = renderNode(p.Changes, p.TargetID);
                    el.replaceWith(newEl);
                    break;

                case "remove":
                    el.remove();
                    break;

                case "add-child":
                    const index = el.children.length;
                    const newChild = renderNode(p.Changes, `${p.TargetID}/${index}`);
                    el.appendChild(newChild);
                    break;
            }
        });
    }

    return {
        mount,
        patch,
    };
})();

function checkLoop() {
    if (window.GovinciWASM.IsDirty()) {
        console.log("dom is dirty")
        const patch = window.GovinciWASM.RenderAgain();
        Govinci.patch(patch);
    }
    requestAnimationFrame(checkLoop);
}

checkLoop(); // inicia o loop
