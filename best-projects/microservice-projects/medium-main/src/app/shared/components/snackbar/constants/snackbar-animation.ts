import { AnimationTriggerMetadata, animate, style, transition, trigger } from '@angular/animations';

export const SnackbarAnimation: AnimationTriggerMetadata = trigger('slideIn', [
  transition(':enter', [
    style({ transform: 'translateY(50%)', opacity: 0 }),
    animate('350ms ease-out', style({ transform: 'translateY(0)', opacity: 1 })),
  ]),
]);
