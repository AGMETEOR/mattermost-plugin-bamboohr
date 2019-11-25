import {id as pluginId} from './manifest';

export default class Plugin {
    // eslint-disable-next-line no-unused-vars
    initialize(registry, store) {}
}

window.registerPlugin(pluginId, new Plugin());
