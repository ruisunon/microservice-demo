import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input } from '@angular/core';

@Component({
  selector: 'app-error-messages',
  standalone: true,
  imports: [CommonModule],
  template: `
    <ul>
      <li *ngFor="let error of errorMessages" class="error">
        {{ error }}
      </li>
    </ul>
  `,
  styleUrls: ['./error-messages.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ErrorMessagesComponent {
  @Input() errorMessages!: string[];
}
