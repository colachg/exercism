class Clock:
    def __init__(self, hour, minute):
        self.minutes = (minute + 60 * hour) % (24 * 60)

    def __repr__(self):
        return "%02d:%02d" % (self.minutes // 60, self.minutes % 60)

    def __eq__(self, other):
        return repr(self) == repr(other)

    def __add__(self, minutes):
        self.minutes = (self.minutes + minutes) % (24 * 60)
        return self

    def __sub__(self, minutes):
        self.minutes = (self.minutes - minutes) % (24 * 60)
        return self
