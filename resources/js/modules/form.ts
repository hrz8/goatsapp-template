import type { Signal } from "signal-polyfill";

export interface InputState  {
  el: HTMLElement | null;
  wrapper?: HTMLElement | null;
  optional?: boolean;
  state: Signal.State<string>;
  regExp: RegExp;
  err: string;
};

export function validateInputForm(...inputs: InputState[]): boolean {
  const success = [];

  for (const input of inputs) {
    if (!input.el) {
      success.push(true);
      continue;
    }
  
    const wrapper = input.wrapper ? input.wrapper : input.el.parentElement;
    const errMsg = wrapper?.querySelector('.error-msg');

    if (!wrapper || !errMsg) {
      success.push(true);
      continue;
    }

    if (input.state.get() === "") {
      success.push(input.optional);
      wrapper.classList.remove('has-error');
      continue;
    }
  
    if (!input.regExp.test(input.state.get())) {
      wrapper.classList.add('has-error');
      if (errMsg && errMsg.tagName.toLowerCase() === 'p' && errMsg.classList.contains('error-msg')) {
        errMsg.textContent = input.err;
        success.push(false);
        continue;
      }

      success.push(true);
    } else {
      wrapper.classList.remove('has-error');
      success.push(true);
    }
  }

  return success.every(s => s === true);
}
