import { Signal } from 'signal-polyfill';
import { effect } from '../../modules/state';
import { validateInputForm } from '../../modules/form';
import type { InputState } from '../../modules/form';

const projectName: InputState = {
  el: document.getElementById('project-name'),
  state: new Signal.State(""),
  regExp: /[\w\s\p{Emoji}]{5,50}/u,
  err: "Invalid value (range: 5-50 chars)",
};

const projectAlias: InputState = {
  el: document.getElementById('project-alias'),
  state: new Signal.State(""),
  regExp: /^[a-z-]{5,25}$/,
  err: "Invalid value (no spaces, only accept dash (-), range: 5-25)",
};

const description: InputState = {
  el: document.getElementById('project-description'),
  state: new Signal.State(""),
  regExp: /.{10,}/,
  optional: true,
  err: "Invalid value (min: 10 chars)",
}

const webhookUrl: InputState = {
  el: document.getElementById('webhook-url'),
  wrapper: document.getElementById('webhook-url')?.parentElement?.parentElement,
  state: new Signal.State(""),
  regExp: /^(https?:\/\/)?(?:localhost|\b(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}\b)(?::\d{1,5})?(?:\/\S*)?$/,
  err: "Invalid url value",
}

if (projectName.el) {
  projectName.el.addEventListener('change', (event) => {
    projectName.state.set((event.target as HTMLInputElement).value);
  });
}

if (projectAlias.el) {
  projectAlias.el.addEventListener('change', (event) => {
    projectAlias.state.set((event.target as HTMLInputElement).value);
  });
}

if (description.el) {
  description.el.addEventListener('change', (event) => {
    description.state.set((event.target as HTMLInputElement).value);
  });
}

if (webhookUrl.el) {
  webhookUrl.el.addEventListener('change', (event) => {
    webhookUrl.state.set((event.target as HTMLInputElement).value);
  });
}

effect(() => {
  const isInputValid = validateInputForm(projectName, projectAlias, description, webhookUrl);
  document
    .querySelectorAll('.form-save')
    .forEach((el) => {
      if (isInputValid) {
        el.removeAttribute('disabled');
        el.classList.remove('cursor-not-allowed');
      } else {
        el.setAttribute('disabled', 'true');
        el.classList.add('cursor-not-allowed');
      }
    });
});

// -- form handler
interface HtmxAfterRequestEvent {
  xhr: XMLHttpRequest;
}

const appNewForm = document.getElementById('project-new-form');

if (appNewForm) {
  appNewForm.addEventListener('htmx:afterRequest', ((event: CustomEvent<HtmxAfterRequestEvent>) => {
    document
      .querySelectorAll('.form-save')
      .forEach((el) => el.classList.remove('hidden'));

    if (event.detail.xhr.status === 200) {
      queueMicrotask(() => {
        document.getElementById('main-toast')?.classList.remove('hidden');
        document.getElementById('main-toast')?.classList.add('flex');
      });

      // reset form
      (projectName.el as HTMLInputElement).value = '';
      (projectAlias.el as HTMLInputElement).value = '';
      (description.el as HTMLInputElement).value = '';
      (webhookUrl.el as HTMLInputElement).value = '';
    }
  }) as EventListener);
  
  appNewForm.addEventListener('htmx:beforeRequest', ((event: CustomEvent<HtmxAfterRequestEvent>) => {
    document
      .querySelectorAll('.form-save')
      .forEach((el) => el.classList.add('hidden'));
  }) as EventListener);
}
