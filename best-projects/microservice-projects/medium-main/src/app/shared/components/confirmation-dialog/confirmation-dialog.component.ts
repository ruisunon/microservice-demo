import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { ConfirmationDialogData } from '@core/models/confirmation-dialog-data.model';

const ConfirmationDialogImports: Array<any> = [CommonModule, MatButtonModule, MatDialogModule];

@Component({
  selector: 'app-confirmation-dialog',
  standalone: true,
  imports: ConfirmationDialogImports,
  template: `
    <h4 mat-dialog-title>Confirmation</h4>
    <div class="px-3">{{ dialogData.label }}</div>

    <div mat-dialog-actions class="buttons">
      <button [mat-dialog-close]="true" mat-button color="primary" class="w-full">Yes</button>
      <button [mat-dialog-close]="false" mat-button class="w-full">No</button>
    </div>
  `,
  styleUrls: ['./confirmation-dialog.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfirmationDialogComponent {
  constructor(@Inject(MAT_DIALOG_DATA) public dialogData: ConfirmationDialogData) {}
}
