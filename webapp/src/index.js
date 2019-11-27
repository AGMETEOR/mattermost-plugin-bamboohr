import {id as pluginId} from './manifest';
import reducer from './reducers';
import Icon from './components/icon';
import {openBambooModal} from './actions';

export default class Plugin {
    async initialize(registry, store) {
        registry.registerReducer(reducer);

        registry.registerMainMenuAction(
            'Bamboo',
            () => store.dispatch(openBambooModal()),
            Icon,
        );
    }
}

window.registerPlugin(pluginId, new Plugin());
