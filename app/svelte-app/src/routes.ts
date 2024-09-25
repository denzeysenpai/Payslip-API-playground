import Calendar from './routes/Calendar/Calendar.svelte';
import Home from './routes/Dashboard/Dashboard.svelte';
import Manage from './routes/Manage/Manage.svelte';
import NotFound from './routes/NotFound/NotFound.svelte';
import Payslip from './routes/Payslip/Payslip.svelte';

export default {
    '/': Home,
    '/dashboard': Home,
    '/calendar': Calendar,
    '/payslip': Payslip,
    '/manage': Manage,
    '*': NotFound
};
