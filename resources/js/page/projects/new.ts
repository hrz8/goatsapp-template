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
      })
    }
  }) as EventListener);
  
  appNewForm.addEventListener('htmx:beforeRequest', ((event: CustomEvent<HtmxAfterRequestEvent>) => {
    document
      .querySelectorAll('.form-save')
      .forEach((el) => el.classList.add('hidden'));

  }) as EventListener);
}
