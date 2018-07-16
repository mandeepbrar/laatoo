import React from 'react';
import sanitizeHtml from 'sanitize-html';

function createHtml(text, sanitize) {
  let retval = text
  if(sanitize) {
    retval = sanitizeHtml(text)
  }
  return {__html: retval}
}

const Html =(props) => {
  return (
    <div className={props.className} style={props.style} dangerouslySetInnerHTML={createHtml(props.children, props.sanitize)}>
    </div>
  )
}

export {Html as Html}
