/*
 *    Copyright (C) 2014 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

// beehive's central module system.
package modules

// Modules provide events, which are described in an EventDescriptor.
type EventDescriptor struct {
	Namespace   string
	Name        string
	Description string
	Options     []PlaceholderDescriptor
}

// Modules offer actions, which are described in an ActionDescriptor.
type ActionDescriptor struct {
	Namespace   string
	Name        string
	Description string
	Options     []PlaceholderDescriptor
}

// A PlaceholderDescriptor shows which in & out values a module expects and returns.
type PlaceholderDescriptor struct {
	Name        string
	Description string
	Type        string
}

// Returns the ActionDescriptor matching an action.
func GetActionDescriptor(action *Action) ActionDescriptor {
	mod := (*GetModule(action.Namespace))
	for _, ac := range mod.Actions() {
		if ac.Name == action.Name {
			return ac
		}
	}

	return ActionDescriptor{}
}

// Returns the EventDescriptor matching an event.
func GetEventDescriptor(event *Event) EventDescriptor {
	mod := (*GetModule(event.Namespace))
	for _, ev := range mod.Events() {
		if ev.Name == event.Name {
			return ev
		}
	}

	return EventDescriptor{}
}
