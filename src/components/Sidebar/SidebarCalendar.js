import React, { useState, useMemo } from 'react';

// SidebarCalendar: small month view with prev/next navigation and today highlight
export default function SidebarCalendar() {
  // displayedMonth: Date object set to 1st of month
  const [displayed, setDisplayed] = useState(() => {
    const now = new Date();
    return new Date(now.getFullYear(), now.getMonth(), 1);
  });

  const today = useMemo(() => {
    const n = new Date();
    return { year: n.getFullYear(), month: n.getMonth(), date: n.getDate() };
  }, []);

  const monthLabel = useMemo(() => {
    // format as "noviembre 2025" (no 'de') and lowercase month as in the mock
    const monthName = displayed.toLocaleString('es-ES', { month: 'long' });
    return `${monthName} ${displayed.getFullYear()}`;
  }, [displayed]);

  const prevMonth = () => {
    setDisplayed((d) => new Date(d.getFullYear(), d.getMonth() - 1, 1));
  };
  const nextMonth = () => {
    setDisplayed((d) => new Date(d.getFullYear(), d.getMonth() + 1, 1));
  };

  // Build matrix of weeks starting Monday (0) to Sunday (6)
  const weeks = useMemo(() => {
    const year = displayed.getFullYear();
    const month = displayed.getMonth();
    const first = new Date(year, month, 1);
    const last = new Date(year, month + 1, 0); // last day of month
    // dayOfWeek Monday-based index
    const mapDay = (jsDay) => (jsDay + 6) % 7;
    const firstIndex = mapDay(first.getDay());
    const numDays = last.getDate();
    const cells = [];
    // fill leading blanks
    for (let i = 0; i < firstIndex; i++) cells.push(null);
    for (let d = 1; d <= numDays; d++) cells.push(new Date(year, month, d));
    // pad to full weeks
    while (cells.length % 7 !== 0) cells.push(null);
    const weeksArr = [];
    for (let i = 0; i < cells.length; i += 7) weeksArr.push(cells.slice(i, i + 7));
    return weeksArr;
  }, [displayed]);

  return (
    <div className="sidebar-calendar-card" style={{ marginTop: 0 }}>
      <div className="card-body p-3">
        <h6 id="instance-49-header" className="card-title text-uppercase bg-white p-2 font-weight-bold" style={{ color: 'var(--moodle-green)', borderRadius: 12, boxShadow: '0 4px 6px rgba(0,0,0,0.1), 0 1px 3px rgba(0,0,0,0.08)', margin: 0 }}>Calendario</h6>

        <div className="card-text content mt-3">
          <div data-region="calendar" className="maincalendar">
            <div className="calendarwrapper" style={{ paddingTop: 8 }}>

              <div style={{ marginTop: 6, background: '#fff', borderRadius: 8, padding: 8, border: '1px solid #e6e6e6' }}>
                <div className="calendar-nav" style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', padding: '4px 6px 8px' }}>
                  <button type="button" aria-label="Mes anterior" onClick={prevMonth} className="calendar-nav-btn" style={{ border: 'none', background: 'transparent', cursor: 'pointer' }}>
                    <span className="calendar-nav-arrow" aria-hidden="true">◄</span>
                  </button>
                  <h4 className="calendar-month-label" style={{ margin: 0, fontWeight: 400, color: '#111', textAlign: 'center' }}>{monthLabel}</h4>
                  <button type="button" aria-label="Mes siguiente" onClick={nextMonth} className="calendar-nav-btn" style={{ border: 'none', background: 'transparent', cursor: 'pointer' }}>
                    <span className="calendar-nav-arrow" aria-hidden="true">►</span>
                  </button>
                </div>
                <table className="calendarmonth" style={{ width: '100%', borderCollapse: 'collapse' }}>
                  <thead>
                    <tr>
                      <th>Lun</th>
                      <th>Mar</th>
                      <th>Mié</th>
                      <th>Jue</th>
                      <th>Vie</th>
                      <th>Sáb</th>
                      <th>Dom</th>
                    </tr>
                  </thead>
                  <tbody>
                    {weeks.map((week, wi) => (
                      <tr key={wi} data-region="month-view-week">
                        {week.map((d, di) => {
                          if (!d) return <td key={di} className="dayblank">&nbsp;</td>;
                          const isToday = d.getFullYear() === today.year && d.getMonth() === today.month && d.getDate() === today.date;
                          return (
                            <td key={di} className={`day ${isToday ? 'today' : ''}`} data-day={d.getDate()} style={{ textAlign: 'center' }}>
                              <div className="d-none d-md-block hidden-phone text-xs-center">
                                <span aria-hidden="true"><span className="day-number-circle" style={isToday ? { background: 'var(--moodle-green)', color: '#fff', display: 'inline-block', width: 28, height: 28, borderRadius: '50%', lineHeight: '28px' } : {}}><span className="day-number">{d.getDate()}</span></span></span>
                              </div>
                            </td>
                          );
                        })}
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>

              <div className="footer" style={{ paddingTop: 0, marginTop: 10 }}>
                <div className="bottom" style={{ display: 'flex', flexDirection: 'column', gap: 6}}>
                  <a href="#" style={{ fontSize: 13 }}>Calendario completo</a>
                  <a href="#" style={{ fontSize: 13 }}>Importar o exportar calendarios</a>
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
