class Clock:
    def __init__(self, hour, minute):
        self.hour = (hour + minute // 60) % 24
        self.minute = minute % 60

    def __repr__(self):
        return "%02d:%02d" % (self.hour % 24, self.minute % 60)

    def __eq__(self, other):
        if other.hour == self.hour and other.minute == self.minute:
            return True
        return False

    def __add__(self, minutes):
        self.hour = self.hour + (self.minute + minutes) // 60
        self.minute = (self.minute + minutes) % 60
        return "%02d:%02d" % (self.hour % 24, self.minute % 60)

    def __sub__(self, minutes):
        if minutes <= 60:
            if minutes <= self.minute:
                self.minute = self.minute - minutes
            else:
                self.minute = self.minute - minutes + 60
                self.hour -= 1
        else:
            self.hour += (self.minute - minutes) // 60
            self.minute = (self.minute - minutes) % 60
        self.hour += 24 if self.hour < 0 else 0
        return "%02d:%02d" % (self.hour % 24, self.minute % 60)
