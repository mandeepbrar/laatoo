import React from 'react';

// Signed-in user context
const SelectableViewContext = React.createContext({
    register: (id, data) => {},
    setItemSelection: (id, data, selection) => {},
    isSelected: (id) => {}
})

export {
    SelectableViewContext
}