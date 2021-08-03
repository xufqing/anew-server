import { useState } from 'react';

export default () => {
    const [consoleWin, setConsoleWin] = useState<any>(null);

    return { consoleWin, setConsoleWin };
};