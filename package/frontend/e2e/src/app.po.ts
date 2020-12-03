import { browser, by, element } from 'protractor';
import { promise as wdpromise, promise } from 'selenium-webdriver';

export class AppPage {
    navigateTo(): wdpromise.Promise<{}> {
        return browser.get('/');
    }

    getPageTitle(): promise.Promise<string> {
        return element(by.css('ion-title')).getText();
    }
}
