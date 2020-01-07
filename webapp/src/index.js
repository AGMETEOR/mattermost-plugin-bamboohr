import {id as pluginId} from './manifest';
import reducer from './reducers';
import Icon from './components/icon';
import BambooAction from './components/bamboo_action';
import {openBambooModal} from './actions';

export default class Plugin {
    async initialize(registry, store) {
        registry.registerReducer(reducer);

        registry.registerMainMenuAction(
            'Bamboo',
            () => store.dispatch(openBambooModal()),
            Icon,
        );
        registry.registerPopoverUserActionsComponent(BambooAction);
    }
}

window.registerPlugin(pluginId, new Plugin());
