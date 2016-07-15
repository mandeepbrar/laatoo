import t from 'tcomb-form';
import React from 'react';
import TinyMCEInput from 'react-tinymce-input';

const TINYMCE_CONFIG = {
  'language'  : 'en',
  'theme'     : 'modern',
  'toolbar'   : 'bold italic underline strikethrough hr | bullist numlist | link unlink | undo redo | alignleft aligncenter alignright | spellchecker code',
  'menubar'   : false,
  'statusbar' : true,
  'resize'    : true,
  'plugins'   : 'link,spellchecker,paste',
  'theme_modern_toolbar_location' : 'top',
  'theme_modern_toolbar_align': 'left'
};

class TextEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      let data = locals.value
      if(!data) {
        data = '';
      }
      let mceconfig = Object.assign({},TINYMCE_CONFIG, locals.config.mceconfig)
      return (
        <div>
          {(locals.config && locals.config.hideLabel)? null :<label>{locals.label}</label>}
          <TinyMCEInput  style={locals.config.style} className={locals.config.className} value={data} rows={locals.config.rows}
            textAreaProps={locals.config.textAreaProps} tinymceConfig={mceconfig} onChange={locals.onChange} />
        </div>
      );
    };
  }
}


export {
  TextEdit as TextEdit
};
