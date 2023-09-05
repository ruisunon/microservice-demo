import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from 'src/app/app.component';
import { appConfig } from 'src/app/app.config';

// eslint-disable-next-line no-console
bootstrapApplication(AppComponent, appConfig).catch((err) => console.error(err));
