import t from 'tcomb-form';
import React from 'react';
import ReactPlayer from 'react-player';

class VideoEdit extends t.form.Component { // extend the base class
  getTemplate() {
    return (locals) => {
      let hideLabel = locals.config && locals.config.hideLabel
      return (
        <div>
          {
            hideLabel?
            null:
            <label>{locals.label}</label>
          }          
          <t.form.Form type={t.Str} value={locals.value} onChange={locals.onChange}/>
          <ReactPlayer url={locals.value} />
        </div>
      );
    };
  }
}


export {
  VideoEdit as VideoEdit
};
