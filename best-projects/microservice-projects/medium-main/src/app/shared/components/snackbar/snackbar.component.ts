import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { SnackbarAnimation } from '@shared/components/snackbar/constants/snackbar-animation';
import { SNACKBAR_CONFIG } from '@shared/components/snackbar/constants/snackbar-config-token';
import { SnackbarConfig } from '@shared/components/snackbar/models/snackbar-config.model';
import { SnackbarType } from '@shared/components/snackbar/models/snackbar-type.model';
import { SnackbarService } from '@shared/components/snackbar/services/snackbar.service';

const SnackbarImports: Array<any> = [CommonModule, MatButtonModule];

@Component({
  selector: 'app-snackbar',
  standalone: true,
  imports: SnackbarImports,
  templateUrl: './snackbar.component.html',
  styleUrls: ['./snackbar.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
  animations: [SnackbarAnimation],
})
export class SnackbarComponent {
  public message: string;
  public title?: string;
  public position: string;
  public type: SnackbarType;
  public buttonLabel: string;

  constructor(@Inject(SNACKBAR_CONFIG) private readonly config: SnackbarConfig, public readonly snackbarService: SnackbarService) {
    this.message = this.config.message;
    this.title = this.config.title;
    this.type = this.config.type ?? 'info';
    this.buttonLabel = this.config.buttonLabel ?? 'Close';
    this.position = this.snackbarService.setSnackbarPosition(this.config.position);
  }
}
